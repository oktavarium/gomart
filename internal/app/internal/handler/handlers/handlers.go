package handlers

import (
	authenticatorer "github.com/oktavarium/gomart/internal/app/internal/authenticatorer"
	"github.com/oktavarium/gomart/internal/app/internal/logger"
	"github.com/oktavarium/gomart/internal/app/internal/orderer"
)

type Handlers struct {
	logger          logger.Logger
	authenticatorer authenticatorer.Authenticatorer
	orderer         orderer.Orderer
}

func NewHandlers(l logger.Logger, a authenticatorer.Authenticatorer, o orderer.Orderer) *Handlers {
	return &Handlers{
		logger:          l,
		authenticatorer: a,
		orderer:         o,
	}
}
