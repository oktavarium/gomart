package orderer

import (
	"context"

	"github.com/oktavarium/gomart/internal/app/internal/model"
)

type Orderer interface {
	MakeOrder(context.Context, string, string) error
	GetOrders(context.Context, string) ([]model.Order, error)
	GetBalance(context.Context, string) (model.Balance, error)
	Withdraw(context.Context, string, string, float32) error
	GetWithdrawals(context.Context, string) ([]model.Withdrawals, error)
}
