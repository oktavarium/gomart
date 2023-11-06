package orders

import (
	"context"

	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
)

type Storage interface {
	UserByOrder(context.Context, string) (string, error)
	NewOrder(context.Context, string, string) error
	UpdateOrder(context.Context, string, string, *int) error
	Orders(context.Context, string) ([]model.Order, error)
	Balance(context.Context, string) (model.Balance, error)
	Withdraw(context.Context, string, string, int) error
	Withdrawals(context.Context, string) ([]model.Withdrawals, error)
}
