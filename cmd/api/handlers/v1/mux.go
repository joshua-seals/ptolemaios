package v1

import (
	"log/slog"

	"github.com/alexedwards/scs/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

// This is the managing mux structure
// All core aspects necessary to instantiate it
// should be provided below.
type Mux struct {
	logger         *slog.Logger
	sessionManager *scs.SessionManager
	db             *pgxpool.Pool
}

func New(logger *slog.Logger, sessionManager *scs.SessionManager, db *pgxpool.Pool) *Mux {
	return &Mux{
		logger:         logger,
		sessionManager: sessionManager,
		db:             db,
	}
}
