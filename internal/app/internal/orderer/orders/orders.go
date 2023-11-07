package orders

import (
	"context"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
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

	return &Orders{
		ctx:      ctx,
		logger:   logger,
		storage:  storage,
		ordersCh: ordersCh,
	}
}

func (o *Orders) OrdersChan() <-chan string {
	return o.ordersCh
}
