package handlers

import "github.com/oktavarium/gomart/internal/app/internal/server/internal/storage"

type Handlers struct {
	storage     storage.Storage
	accrualAddr string
}

func NewHandlers(s storage.Storage, accuralAddr string) *Handlers {
	return &Handlers{s, accuralAddr}
}
