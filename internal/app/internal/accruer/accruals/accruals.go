package accruals

import (
	"context"
	"time"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
	"github.com/oktavarium/gomart/internal/app/internal/storager"
)

var accrualPath = "/api/orders"
var defaultBufferize uint = 10
var defaultRequestInterval = 1 * time.Second

type Accruals struct {
	ctx         context.Context
	logger      logger.Logger
	accrualAddr string
	storage     storager.Storager
}

func NewAccruals(
	ctx context.Context,
	logger logger.Logger,
	accrualAddr string,
	storage storager.Storager,
	ordersCh <-chan string,
	bufferSize uint,
) (*Accruals, error) {

	accruals := &Accruals{
		ctx:         ctx,
		logger:      logger,
		accrualAddr: accrualAddr,
		storage:     storage,
	}

	unproccessedOrdersCh := accruals.startLoader(ctx, bufferSize)
	pointsCh := accruals.startExecutor(fanIn(ordersCh, unproccessedOrdersCh), bufferSize)

	accruals.startUpdater(pointsCh)

	return accruals, nil
}
