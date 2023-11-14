package provider

import (
	"context"
	"fmt"

	"github.com/oktavarium/gomart/internal/app/internal/accruer"
	"github.com/oktavarium/gomart/internal/app/internal/accruer/accruals"
	"github.com/oktavarium/gomart/internal/app/internal/authenticatorer"
	"github.com/oktavarium/gomart/internal/app/internal/authenticatorer/authenticator"
	"github.com/oktavarium/gomart/internal/app/internal/configer"
	"github.com/oktavarium/gomart/internal/app/internal/configer/config"
	"github.com/oktavarium/gomart/internal/app/internal/handler"
	"github.com/oktavarium/gomart/internal/app/internal/handler/handlers"
	"github.com/oktavarium/gomart/internal/app/internal/logger"
	"github.com/oktavarium/gomart/internal/app/internal/logger/log"
	"github.com/oktavarium/gomart/internal/app/internal/orderer"
	"github.com/oktavarium/gomart/internal/app/internal/orderer/orders"
	"github.com/oktavarium/gomart/internal/app/internal/router"
	"github.com/oktavarium/gomart/internal/app/internal/router/chirouter"
	"github.com/oktavarium/gomart/internal/app/internal/storager"
	"github.com/oktavarium/gomart/internal/app/internal/storager/storage/pg"
)

type ServiceProvider struct {
	configer      configer.Configer
	logger        logger.Logger
	storager      storager.Storager
	authenticator authenticatorer.Authenticatorer
	orderer       orderer.Orderer
	accruer       accruer.Accruer
	handler       handler.Handler
	router        router.Router
}

func NewServiceProvider(ctx context.Context) (*ServiceProvider, error) {
	sp := new(ServiceProvider)
	var err error

	sp.configer, err = config.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("error on creating config: %w", err)
	}

	sp.logger = log.NewLogger(sp.configer.LogLevel())
	sp.storager, err = pg.NewStorage(ctx, sp.logger, sp.configer.DatabaseURI())
	if err != nil {
		return nil, fmt.Errorf("error on creating storage: %w", err)
	}

	sp.authenticator = authenticator.NewAuthenticator(sp.logger, sp.configer.DatabaseURI(), sp.storager)
	sp.orderer = orders.NewOrders(ctx, sp.logger, sp.storager, sp.configer.BufferSize())
	sp.accruer, err = accruals.NewAccruals(
		ctx,
		sp.logger,
		sp.configer.AccrualAddress(),
		sp.storager,
		sp.orderer.OrdersChan(),
		sp.configer.BufferSize(),
	)
	if err != nil {
		return nil, fmt.Errorf("error on creating accruer: %w", err)
	}

	sp.handler = handlers.NewHandlers(sp.logger, sp.authenticator, sp.orderer)
	sp.router = chirouter.NewRouter(ctx, sp.logger, sp.configer.Address(), sp.handler)

	return sp, nil
}

func (sp *ServiceProvider) Run() error {
	return sp.router.Run()
}
