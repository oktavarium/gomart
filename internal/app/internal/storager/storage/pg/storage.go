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
	users  map[string]string
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
		users:  make(map[string]string),
	}

	if err := s.cacheUsers(ctx); err != nil {
		return nil, fmt.Errorf("error on caching users: %w", err)
	}

	return s, nil
}
