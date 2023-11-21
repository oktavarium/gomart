package pg

import (
	"context"
	"fmt"
)

func (s *storage) WithdrawInTx(ctx context.Context, user, order string, sum float32) func() error {
	return func() error {
		return s.Withdraw(ctx, user, order, sum)
	}
}

func (s *storage) UpdateBalanceInTx(ctx context.Context, user string, sum float32) func() error {
	return func() error {
		return s.UpdateBalance(ctx, user, sum)
	}
}

func (s *storage) MakeInTx(ctx context.Context, funcs ...func() error) error {
	tx, err := s.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error on begin tx: %w", err)
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			s.logger.Error(err)
		}
	}()

	for _, fn := range funcs {
		if err := fn(); err != nil {
			return fmt.Errorf("error on making func in tx: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error on tx commit: %w", err)
	}

	return nil
}
