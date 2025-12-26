package postgres

import (
	"destinations-suggester/internal/pkg/sl"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

const migrationsPath = "file://migrations/postgres"

func ApplyMigrations(connString ConnString) error {
	conf, err := pgx.ParseConfig(string(connString))
	if err != nil {
		return fmt.Errorf("cannot parse conn string: %w", err)
	}

	m, err := migrate.New(migrationsPath, conf.ConnString())
	if err != nil {
		return fmt.Errorf("cannot create migrator: %w", err)
	}

	defer func() {
		sErr, dbErr := m.Close()
		if sErr != nil {
			slog.Error("closing go migrate source error", sl.Error(err))
		}
		if dbErr != nil {
			slog.Error("closing go migrate database error", sl.Error(err))
		}
	}()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("cannot apply migrations failed: %w", err)
	}

	return nil
}
