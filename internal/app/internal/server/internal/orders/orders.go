package orders

import "github.com/oktavarium/gomart/internal/app/internal/server/internal/orders/internal/accruals"

var defaultBufferize uint = 10

type Orders struct {
	storage  Storage
	accruals *accruals.Accruals
	ordersCh chan string
}

func NewOrders(accrualAddr string, storage Storage, bufferSize uint) *Orders {
	if bufferSize == 0 {
		bufferSize = defaultBufferize
	}

	ordersCh := make(chan string, bufferSize)

	return &Orders{
		storage:  storage,
		ordersCh: ordersCh,
		accruals: accruals.NewAccruals(accrualAddr, storage, ordersCh, bufferSize),
	}
}
