package core

import (
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Initialize a new session manager.
// Adjust the Lifetime here to create
// shorter or longer sessions as needed.
// Currently this is hardcoded to keep it simple.
func NewSessionManager(db *pgxpool.Pool) *scs.SessionManager {
	sessionManager := scs.New()
	sessionManager.Store = pgxstore.New(db)
	sessionManager.Lifetime = 1 * time.Hour
	return sessionManager
}
