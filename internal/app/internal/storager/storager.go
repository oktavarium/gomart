package storager

import (
	"context"

	"github.com/oktavarium/gomart/internal/app/internal/model"
)

type Storager interface {
	GetUserByOrder(context.Context, string) (string, error)
	MakeOrder(context.Context, string, string) error
	UpdateOrder(context.Context, string, string, float32) error
	GetOrders(context.Context, string) ([]model.Order, error)
	GetOrdersByStatus(context.Context, []string) ([]string, error)
	GetBalance(context.Context, string) (float32, float32, error)
	Withdraw(context.Context, string, string, float32) error
	UpdateBalance(context.Context, string, float32) error
	GetWithdrawals(context.Context, string) ([]model.Withdrawals, error)
	UserExists(string) bool
	RegisterUser(context.Context, string, string, string) error
	GetUserHashAndSalt(context.Context, string) (string, string, error)

	WithdrawInTx(context.Context, string, string, float32) func() error
	UpdateBalanceInTx(context.Context, string, float32) func() error
	MakeInTx(context.Context, ...func() error) error
}
