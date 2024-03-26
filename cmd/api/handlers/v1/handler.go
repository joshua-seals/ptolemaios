package v1

import (
	"log/slog"

	"github.com/alexedwards/scs/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CoreHandler struct {
	logger         *slog.Logger
	sessionManager *scs.SessionManager
	db             *pgxpool.Pool
}

func New(logger *slog.Logger, sessionManager *scs.SessionManager, db *pgxpool.Pool) *CoreHandler {
	return &CoreHandler{
		logger:         logger,
		sessionManager: sessionManager,
		db:             db,
	}
}
