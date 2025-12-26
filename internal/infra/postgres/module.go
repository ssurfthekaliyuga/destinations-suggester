package postgres

import (
	"context"
	"destinations-suggester/internal/infra/postgres/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
	"log/slog"
)

var Module = fx.Module("postgres",
	fx.Provide(
		Connect,
		ApplyMigrations,
		repositories.NewPlaces,
		repositories.NewPlaces,
	),
	fx.Invoke(
		func(lc fx.Lifecycle, pool *pgxpool.Pool) {
			lc.Append(fx.Hook{
				OnStop: func(ctx context.Context) error {
					slog.Info("closing postgres connection")
					pool.Close()
					return nil
				},
			})
		},
	),
)
