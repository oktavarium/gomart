package orders

import (
	"context"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
	accruals "github.com/oktavarium/gomart/internal/app/internal/orderer/orders/internal/accruer"
	"github.com/oktavarium/gomart/internal/app/internal/storager"
)

var defaultBufferize uint = 10

type Orders struct {
	ctx      context.Context
	logger   logger.Logger
	storage  storager.Storager
	ordersCh chan string
}

func NewOrders(
	ctx context.Context,
	logger logger.Logger,
	accrualAddress string,
	storage storager.Storager,
	bufferSize uint,
) *Orders {

	if bufferSize == 0 {
		bufferSize = defaultBufferize
	}

	ordersCh := make(chan string, bufferSize)

	go func() {
		<-ctx.Done()
		close(ordersCh)
	}()

	accruals.NewAccruals(
		ctx,
		logger,
		accrualAddress,
		storage,
		ordersCh,
		bufferSize,
	)

	return &Orders{
		ctx:      ctx,
		logger:   logger,
		storage:  storage,
		ordersCh: ordersCh,
	}
}
