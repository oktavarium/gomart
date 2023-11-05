package orders

import (
	"context"

	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
)

type Storage interface {
	GetUserByOrder(context.Context, string) (string, error)
	CreateOrder(context.Context, string, string, string) error
	GetOrders(context.Context, string) ([]model.Order, error)
	GetBalance(context.Context, string) (model.Balance, error)
	Withdraw(context.Context, string, string, int) error
	GetWithdrawals(context.Context, string) ([]model.Withdrawals, error)
}
