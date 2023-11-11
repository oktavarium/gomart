package orderer

import (
	"context"

	"github.com/oktavarium/gomart/internal/app/internal/model"
)

type Orderer interface {
	NewOrder(context.Context, string, string) error
	Orders(context.Context, string) ([]model.Order, error)
	Balance(context.Context, string) (model.Balance, error)
	Withdraw(context.Context, string, string, float32) error
	Withdrawals(context.Context, string) ([]model.Withdrawals, error)
	OrdersChan() <-chan string
}
