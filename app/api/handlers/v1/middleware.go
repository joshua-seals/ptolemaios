package v1

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/joshua-seals/ptolemaios/internal/data/models"
)

// Executed before a request hits servemux
// Debug via browser dev tools if issues arise with browsers and accessing
// the site.
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r)
	})
}

func (m *Mux) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.logger.Info("requested endpoint", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (m *Mux) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This function will always be called when the stack
		// is unwound following a panic
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "closed")
				errS := fmt.Errorf("%s", err)
				m.serverErrorResponse(w, r, errS)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// This authenticate middleware supports bearer token auth
func (m *Mux) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ensure that we indicate authorization may vary
		w.Header().Add("Vary", "Authorization")

		// Returns "" empty string if nothing is found.
		authorizationHeader := r.Header.Get("Authorization")

		// If no auth header, set user as anonymous
		if authorizationHeader == "" {
			w.Header().Set("WWW-Authenticate", "Bearer")
			message := "invalid or missing auth token"
			http.Error(w, message, http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			w.Header().Set("WWW-Authenticate", "Bearer")
			message := "invalid or missing auth token"
			http.Error(w, message, http.StatusUnauthorized)
			return
		}

		token := headerParts[1]

		valid, err := models.IsValidAdmin(m.db, token)
		if err != nil {
			m.logger.Error("Error with token validation", err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
		}
		if !valid {
			w.Header().Set("WWW-Authenticate", "Bearer")
			message := "invalid or missing auth token"
			http.Error(w, message, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
