package handlers

import (
	"github.com/oktavarium/gomart/internal/app/internal/authenticator"
	"github.com/oktavarium/gomart/internal/app/internal/logger"
	"github.com/oktavarium/gomart/internal/app/internal/orderer"
)

type Handlers struct {
	logger        logger.Logger
	authenticator authenticator.Authenticator
	orderer       orderer.Orderer
}

func NewHandlers(l logger.Logger, a authenticator.Authenticator, o orderer.Orderer) *Handlers {
	return &Handlers{
		logger:        l,
		authenticator: a,
		orderer:       o,
	}
}
