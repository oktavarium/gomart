package orders

import (
	"context"
	"fmt"

	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/orders/internal/accruals"
)

type Orders struct {
	storage  Storage
	ordersCh chan string
	accruals *accruals.Accruals
}

func NewOrders(storage Storage, accuralAddr string) *Orders {
	accruals := accruals.NewAccruals(accuralAddr)

	orders := &Orders{
		storage:  storage,
		ordersCh: make(chan string, 10),
		accruals: accruals,
	}

	inCh := accruals.NewExecutor(orders.ordersCh, 10)
	go orders.NewOrdersUpdater(inCh)

	return orders
}

func (o *Orders) NewOrdersUpdater(inCh <-chan model.Points) {
	for points := range inCh {
		if points.Status == REGISTERED || points.Status == PROCESSING {
			points.Status = PROCESSING
			err := o.UpdateOrder(context.TODO(), points.Order, points.Status, points.Accrual)
			if err != nil {
				fmt.Println("SOME ERROR ON UPDATING POINTS")
			}
		}
	}
}
