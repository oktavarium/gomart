package orders

import (
	"fmt"

	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
)

type Orders struct {
	storage Storage
}

func NewOrders(storage Storage) *Orders {
	return &Orders{storage}
}

func (o *Orders) NewOrder(user, order string) error {
	order = compressOrderNumber(order)
	if !checkOrderNumber(order) {
		return ErrWrongOrderNum
	}

	dbUser, err := o.storage.GetUserByOrder(order)
	if err != nil {
		return fmt.Errorf("error on getting user by order: %w", err)
	}

	if user == dbUser {
		return ErrLoadedOrder
	}

	if len(dbUser) != 0 {
		return ErrAnotherUserOrder
	}

	err = o.storage.CreateOrder(user, order, string(NEW))
	if err != nil {
		return fmt.Errorf("error creating new order: %w", err)
	}

	return nil
}

func (o *Orders) GetOrders(user string) ([]model.Order, error) {
	orders, err := o.storage.GetOrders(user)
	if err != nil {
		return nil, fmt.Errorf("error on getting orders: %w", err)
	}

	return orders, nil
}
