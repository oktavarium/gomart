package app

import (
	"context"
	"fmt"

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
	"github.com/oktavarium/gomart/internal/app/internal/pointstorer"
	"github.com/oktavarium/gomart/internal/app/internal/pointstorer/pointmock"
	"github.com/oktavarium/gomart/internal/app/internal/pointstorer/pointstore"
	"github.com/oktavarium/gomart/internal/app/internal/router"
	"github.com/oktavarium/gomart/internal/app/internal/router/chirouter"
	"github.com/oktavarium/gomart/internal/app/internal/storager"
	"github.com/oktavarium/gomart/internal/app/internal/storager/storage/pg"
)

type serviceProvider struct {
	configer        configer.Configer
	logger          logger.Logger
	storager        storager.Storager
	authenticatorer authenticatorer.Authenticatorer
	orderer         orderer.Orderer
	pointstore      pointstorer.PointStorer
	handler         handler.Handler
	router          router.Router
}

func newServiceProvider(ctx context.Context) (*serviceProvider, error) {
	sp := new(serviceProvider)
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

	sp.authenticatorer = authenticator.NewAuthenticator(sp.logger, sp.configer.DatabaseURI(), sp.storager)
	if sp.configer.TestMode() {
		sp.logger.Debug("TESTMODE")
		sp.pointstore = pointmock.NewPointmock(sp.logger)
	} else {
		sp.logger.Debug("NOTESTMODE")
		sp.pointstore = pointstore.NewPointStore(sp.logger, sp.configer.AccrualAddress())
	}

	sp.orderer = orders.NewOrders(ctx, sp.logger, sp.pointstore, sp.storager, sp.configer.BufferSize())
	sp.handler = handlers.NewHandlers(sp.logger, sp.authenticatorer, sp.orderer)
	sp.router = chirouter.NewRouter(ctx, sp.logger, sp.configer.Address(), sp.handler)

	return sp, nil
}

func (sp *serviceProvider) Run() error {
	return sp.router.Run()
}
