package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oktavarium/gomart/internal/app/internal/logger"
)

type storage struct {
	*pgxpool.Pool
	logger logger.Logger
}

func NewStorage(ctx context.Context, logger logger.Logger, dbURI string) (*storage, error) {
	if err := makeMigration(dbURI); err != nil {
		return nil, fmt.Errorf("error on db migration: %w", err)
	}

	pool, err := pgxpool.New(ctx, dbURI)
	if err != nil {
		return nil, fmt.Errorf("error on creating pgxpool: %w", err)
	}

	s := &storage{
		Pool:   pool,
		logger: logger,
	}

	return s, nil
}
