package orders

import (
	"context"
	"fmt"

	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
)

type Orders struct {
	storage Storage
}

func NewOrders(storage Storage) *Orders {
	return &Orders{storage}
}

func (o *Orders) NewOrder(ctx context.Context, user, order string) error {
	order = compressOrderNumber(order)
	if !checkOrderNumber(order) {
		return ErrWrongOrderNumber
	}

	dbUser, err := o.storage.GetUserByOrder(ctx, order)
	if err != nil {
		return fmt.Errorf("error on getting user by order: %w", err)
	}

	if user == dbUser {
		return ErrLoadedOrder
	}

	if len(dbUser) != 0 {
		return ErrAnotherUserOrder
	}

	err = o.storage.CreateOrder(ctx, user, order, string(NEW))
	if err != nil {
		return fmt.Errorf("error creating new order: %w", err)
	}

	return nil
}

func (o *Orders) GetOrders(ctx context.Context, user string) ([]model.Order, error) {
	orders, err := o.storage.GetOrders(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error on getting orders: %w", err)
	}

	return orders, nil
}

func (o *Orders) GetBalance(ctx context.Context, user string) (model.Balance, error) {
	balance, err := o.storage.GetBalance(ctx, user)
	if err != nil {
		return balance, fmt.Errorf("error on getting balance: %w", err)
	}

	return balance, nil
}

func (o *Orders) Withdraw(ctx context.Context, user, order string, sum int) error {
	order = compressOrderNumber(order)
	if !checkOrderNumber(order) {
		return ErrWrongOrderNumber
	}

	balance, err := o.storage.GetBalance(ctx, user)
	if err != nil {
		return fmt.Errorf("error on getting balance: %w", err)
	}

	if balance.Current < sum {
		return ErrNotEnoughBalance
	}

	if err := o.storage.Withdraw(ctx, user, order, sum); err != nil {
		return fmt.Errorf("error on withdrawal")
	}

	return nil
}

func (o *Orders) GetWithdrawals(ctx context.Context, user string) ([]model.Withdrawals, error) {
	return nil, nil
}
