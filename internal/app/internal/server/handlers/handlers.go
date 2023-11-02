package handlers

import "github.com/oktavarium/gomart/internal/app/internal/server/storage"

type Handlers struct {
	storage     storage.Storage
	accrualAddr string
}

func NewHandlers(s storage.Storage, accuralAddr string) *Handlers {
	return &Handlers{s, accuralAddr}
}
