package accruals

import (
	"context"
	"time"

	"github.com/oktavarium/gomart/internal/app/internal/model"
	"github.com/oktavarium/gomart/internal/app/internal/storager"
)

var accrualPath = "/api/orders"
var defaultBufferize uint = 10
var defaultRequestInterval = 1 * time.Second

type Accruals struct {
	ctx         context.Context
	accrualAddr string
	storage     storager.Storager
}

func NewAccruals(
	ctx context.Context,
	accrualAddr string,
	storage storager.Storager,
	ordersCh <-chan string,
	bufferSize uint,
) *Accruals {

	accruals := &Accruals{
		ctx:         ctx,
		accrualAddr: accrualAddr,
		storage:     storage,
	}

	pointsCh := accruals.startExecutor(ordersCh, bufferSize)
	accruals.startUpdater(pointsCh)

	return accruals
}

func (a *Accruals) startExecutor(orders <-chan string, bufferSize uint) <-chan model.Points {
	if bufferSize == 0 {
		bufferSize = defaultBufferize
	}

	outCh := make(chan model.Points, bufferSize)

	go func() {
		ticker := time.NewTicker(defaultRequestInterval)
		for range ticker.C {
			select {
			case <-a.ctx.Done():
				close(outCh)
				return
			default:
				if order, ok := <-orders; ok {
					points, err := getPoints(order)
					if err != nil {
						continue
					}
					outCh <- points
				}
			}
		}
	}()

	return outCh
}
