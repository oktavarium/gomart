package storage

import "github.com/oktavarium/gomart/internal/app/internal/server/storage/pg"

func NewPGStorage(dbURI string) (Storage, error) {
	return pg.NewStorage(dbURI)
}
