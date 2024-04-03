package v1

import (
	"fmt"
	"net/http"
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

func (c *Mux) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.logger.Info("requested endpoint", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (c *Mux) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This function will always be called when the stack
		// is unwound following a panic
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "closed")
				errS := fmt.Errorf("%s", err)
				c.serverErrorResponse(w, r, errS)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
