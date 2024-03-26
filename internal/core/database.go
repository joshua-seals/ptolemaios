package core

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenDB(dsn string) (*pgxpool.Pool, error) {

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
