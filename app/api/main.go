package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	v1 "github.com/joshua-seals/ptolemaios/app/api/handlers/v1"
	"github.com/joshua-seals/ptolemaios/internal/core"
)

// func init() {
// 	// We must register the ptolemaios models
// 	// to include in our session data with scs
// 	gob.Register(models.something{})
// }

func main() {
	// Instantiate the config struct
	cfg := core.NewConfig()
	logger := core.NewLogger(os.Stdout)

	// ** Remove this and support database
	// app functionality
	// clone helx-apps
	// err := core.CloneBranch(cfg.AppUrl, cfg.AppBranch)
	// if err != nil {
	// 	logger.Error(err.Error())
	// }

	db, err := core.OpenDB(cfg.Db.Dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// Initialize a new session manager.
	// Adjust the Lifetime here to create
	// shorter or longer sessions as needed.
	sessionManager := core.NewSessionManager(db)

	mux := v1.New(logger, sessionManager, db)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      mux.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.Env)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)

}
