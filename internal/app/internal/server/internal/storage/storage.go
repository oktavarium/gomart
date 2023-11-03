package storage

import (
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/storage/memory"
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/storage/pg"
)

func NewStorage(dbURI string) (Storage, error) {
	if len(dbURI) == 0 {
		return memory.NewStorage(), nil
	}

	return pg.NewStorage(dbURI)
}
