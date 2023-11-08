package pg

import (
	"context"

	"github.com/oktavarium/gomart/internal/app/internal/model"
)

func (s *storage) NewOrder(ctx context.Context, user, number string) error {

	return nil
}

func (s *storage) UpdateOrder(ctx context.Context, number, status string, accrual *int) error {

	return nil
}

func (s *storage) Orders(ctx context.Context, user string) ([]model.Order, error) {

	return orders, nil
}

func (s *storage) OrdersByStatus(ctx context.Context, statuses []string) ([]string, error) {
	return nil, nil
}

func (s *storage) ChekUserOrder(ctx context.Context, user, order string) error {

	return nil
}
