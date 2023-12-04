package accruals

import (
	"context"
	"time"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
	"github.com/oktavarium/gomart/internal/app/internal/pointstorer"
	"github.com/oktavarium/gomart/internal/app/internal/storager"
)

var defaultBufferize uint = 10
var defaultRequestInterval = 1 * time.Second

type Accruals struct {
	ctx     context.Context
	logger  logger.Logger
	ps      pointstorer.PointStorer
	storage storager.Storager
}

func NewAccruals(
	ctx context.Context,
	logger logger.Logger,
	ps pointstorer.PointStorer,
	storage storager.Storager,
	ordersCh <-chan string,
	bufferSize uint,
) *Accruals {

	accruals := &Accruals{
		ctx:     ctx,
		logger:  logger,
		ps:      ps,
		storage: storage,
	}

	unprocessedOrdersCh := accruals.startLoader(ctx, bufferSize)
	pointsCh := accruals.startExecutor(fanIn(ordersCh, unprocessedOrdersCh), bufferSize)

	accruals.startUpdater(pointsCh)

	return accruals
}
