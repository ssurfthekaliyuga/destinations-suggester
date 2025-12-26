package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, connString ConnString) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, string(connString))
	if err != nil {
		return nil, fmt.Errorf("cannot create pgxpool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("cannot ping postgres: %w", err)
	}

	return pool, nil
}
