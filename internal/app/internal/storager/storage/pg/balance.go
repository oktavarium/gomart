package pg

import (
	"context"

	"github.com/oktavarium/gomart/internal/app/internal/model"
)

func (s *storage) Balance(ctx context.Context, user string) (model.Balance, error) {

	return model.Balance(balance), nil
}

func (s *storage) Withdraw(ctx context.Context, user, order string, sum int) error {

	return nil
}

func (s *storage) Withdrawals(ctx context.Context, user string) ([]model.Withdrawals, error) {
	return nil, nil
}
