package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/joshua-seals/ptolemaios/internal/core"
	"github.com/joshua-seals/ptolemaios/internal/data/models"
)

func (m *Mux) createAuthToken(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := m.readJSON(w, r, &input)
	if err != nil {
		m.logger.Error("decode json", err)
		http.Error(w, "Error decoding JSON body", http.StatusBadRequest)
		return
	}
	v := core.NewValidator()

	models.ValidateEmail(v, input.Email)
	models.ValidatePasswordPlaintext(v, input.Password)

	if !v.Valid() {
		m.logger.Error("validation failed")
		http.Error(w, "Authentication failure", http.StatusUnauthorized)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user, err := models.GetByEmail(ctx, m.db, input.Email)

	if err != nil {
		m.logger.Error("User account not found", err)
		http.Error(w, "Authentication failure", http.StatusUnauthorized)
		return
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		m.logger.Error("Account mismatch", err)
		http.Error(w, "Authentication failure", http.StatusUnauthorized)
		return
	}

	if !match {
		m.logger.Error("Account mismatch", err)
		http.Error(w, "Authentication failure", http.StatusUnauthorized)
		return
	}

	// If password is correct, we issue the token
	token, err := models.NewToken(m.db, user.ID, 4*time.Hour, models.ScopeAuthentication)
	if err != nil {
		m.logger.Error("Error from New Token creation", err)
		http.Error(w, "error processing", http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(token)
	if err != nil {
		m.logger.Error("marshal json", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	// Write response to http.ResponseWriter
	js = append(js, '\n')
	w.Header().Set("Content-Type", "mlication/json")
	w.Write(js)
}
