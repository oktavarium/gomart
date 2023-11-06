package storage

import (
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/storage/memory"
)

func NewStorage(dbURI string) (*memory.Storage, error) {
	return memory.NewStorage(), nil
	// if len(dbURI) == 0 {

	// }

	// return pg.NewStorage(dbURI)
}
