package storager

import (
	"context"

	"github.com/oktavarium/gomart/internal/app/internal/model"
)

type Storager interface {
	UserByOrder(context.Context, string) (string, error)
	NewOrder(context.Context, string, string) error
	UpdateOrder(context.Context, string, string, float32) error
	Orders(context.Context, string) ([]model.Order, error)
	OrdersByStatus(context.Context, []string) ([]string, error)
	Balance(context.Context, string) (model.Balance, error)
	Withdraw(context.Context, string, string, float32) error
	Withdrawals(context.Context, string) ([]model.Withdrawals, error)
	UserExists(context.Context, string) (bool, error)
	RegisterUser(context.Context, string, string, string) error
	UserHashAndSalt(context.Context, string) (string, string, error)
}
