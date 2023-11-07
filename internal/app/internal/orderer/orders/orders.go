package orders

import (
	"github.com/oktavarium/gomart/internal/app/internal/logger"
	"github.com/oktavarium/gomart/internal/app/internal/storager"
)

var defaultBufferize uint = 10

type Orders struct {
	logger   logger.Logger
	storage  storager.Storager
	ordersCh chan string
}

func NewOrders(logger logger.Logger, storage storager.Storager, bufferSize uint) *Orders {
	if bufferSize == 0 {
		bufferSize = defaultBufferize
	}

	ordersCh := make(chan string, bufferSize)

	return &Orders{
		logger:   logger,
		storage:  storage,
		ordersCh: ordersCh,
	}
}

func (o *Orders) OrdersChan() <-chan string {
	return o.ordersCh
}
