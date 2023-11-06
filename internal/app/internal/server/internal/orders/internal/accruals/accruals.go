package accruals

import (
	"context"
	"fmt"
	"time"

	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
)

var accrualPath = "/api/orders"
var defaultBufferize uint = 10
var defaultRequesterInterval = 1 * time.Second

type Accruals struct {
	accrualAddr string
	storage     Storage
}

func NewAccruals(accrualAddr string, storage Storage, ordersCh <-chan string, bufferSize uint) *Accruals {
	accruals := &Accruals{
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
		ticker := time.NewTicker(defaultRequesterInterval)
		for range ticker.C {
			if order, ok := <-orders; ok {
				points, err := getPoints(order)
				if err != nil {
					continue
				}
				outCh <- points
			}
		}
	}()

	return outCh
}

func (a *Accruals) startUpdater(pointsCh <-chan model.Points) {
	go func() {
		for points := range pointsCh {
			if points.Status == REGISTERED || points.Status == PROCESSING {
				points.Status = PROCESSING
				err := a.storage.UpdateOrder(context.TODO(), points.Order, points.Status, points.Accrual)
				if err != nil {
					fmt.Println("SOME ERROR ON UPDATING POINTS")
				}
			}
		}
	}()
}
