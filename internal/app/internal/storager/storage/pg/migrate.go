package pg

import (
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var migrateFS embed.FS

func makeMigration(dbURI string) error {
	d, err := iofs.New(migrateFS, "migrations")
	if err != nil {
		return fmt.Errorf("error on creating iofs: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, dbURI)
	if err != nil {
		return fmt.Errorf("error on creating migration: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error on making migration: %w", err)
	}

	return nil
}
