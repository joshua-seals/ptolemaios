package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	v1 "github.com/joshua-seals/ptolemaios/cmd/api/handlers/v1"
	"github.com/joshua-seals/ptolemaios/internal/core"
)

type config struct {
	port    string
	env     string
	version string
	db      struct {
		dsn string
	}
}

type application struct {
	core   *v1.CoreHandler
	config config
}

// func init() {
// 	// We must register the ptolemaios models
// 	// to include in our session data with scs
// 	gob.Register(models.something{})
// }

func main() {
	// Instantiate the config struct
	var cfg config
	cfg.env = os.Getenv("BUILD_REF")
	cfg.port = os.Getenv("APIPORT")
	cfg.version = os.Getenv("VERSION")
	cfg.db.dsn = os.Getenv("DB_DSN")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := core.OpenDB(cfg.db.dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// Initialize a new session manager.
	// Adjust the Lifetime here to create
	// shorter or longer sessions as needed.
	sessionManager := core.NewSessionManager(db)

	core := v1.New(logger, sessionManager, db)
	app := &application{
		core:   core,
		config: cfg,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.port),
		Handler:      app.core.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)

}
