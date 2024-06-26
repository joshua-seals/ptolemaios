package main

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joshua-seals/ptolemaios/app/tooling/cmd"
	"github.com/joshua-seals/ptolemaios/internal/data/schema"
)

// This tool is used to seed the database

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// This allows us to add an initial admin password to the database.
	pwd := os.Getenv("ADMIN_PASSWD")
	// The Darwin lib reqiures that pgx follow database/sql interfaces.
	db_dsn := os.Getenv("DB_DSN")
	db, err := sql.Open("pgx", db_dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1) // Exit if there is a problem with DB connection
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1) // Exit if cannot ping DB
	}
	// Test to ensure db is not already up.
	// This is useful in event of a pod deletion/death
	// since database seeding and migrations are done via
	// the intiContainer which runs for every newly started pod.
	// This action will cause an ERROR originally in the db
	// assuming nothing has been migrated yet.
	stmt := `select ptolemaios_id from ptolemaios where ptolemaios_id=1;`
	_, err = db.Exec(stmt)
	if err == nil {
		// Assume we got the statement back
		os.Exit(0)
	}
	// Perform migrations
	err = schema.Migrate(db)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Migrations Complete")

	// // Seed the database
	// err = schema.Seed(db)
	// if err != nil {
	// 	logger.Error(err.Error())
	// 	os.Exit(1)
	// }
	// logger.Info("Seed Complete")

	// Generate an admin password hash
	// here we use pwd from ADMIN_PASSWD in Makefile.
	hash, err := cmd.CreateAdminPassword(pwd)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	// here we create an annonymous struct
	admin := struct {
		name  string
		email string
		hash  []byte
	}{
		name:  "admin",
		email: "admin@renci.org",
		hash:  hash,
	}

	// Hard coded and in need of some love at a later date.
	adminSeed := `INSERT into users (name, email, password_hash) VALUES
	  ($1,$2,$3);`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = db.ExecContext(ctx, adminSeed, admin.name, admin.email, admin.hash)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
