package storager

import (
	"context"

	"github.com/oktavarium/gomart/internal/app/internal/model"
)

type Storager interface {
	UserByOrder(context.Context, string) (string, error)
	NewOrder(context.Context, string, string) error
	UpdateOrder(context.Context, string, string, *int) error
	Orders(context.Context, string) ([]model.Order, error)
	Balance(context.Context, string) (model.Balance, error)
	Withdraw(context.Context, string, string, int) error
	Withdrawals(context.Context, string) ([]model.Withdrawals, error)
	UserExists(context.Context, string) (bool, error)
	RegisterUser(context.Context, string, string, string) error
	UserHashAndSalt(context.Context, string) (string, string, error)
}
