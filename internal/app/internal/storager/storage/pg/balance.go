package pg

import (
	"context"
	"fmt"

	"github.com/oktavarium/gomart/internal/app/internal/model"
)

func (s *storage) Balance(ctx context.Context, user string) (model.Balance, error) {
	var balance model.Balance
	row := s.QueryRow(ctx, `SELECT balance, withdrawn FROM user WHERE user = $1`, user)
	if err := row.Scan(&balance.Current, &balance.Withdrawn); err != nil {
		return balance, fmt.Errorf("error on selecting values: %w", err)
	}

	return model.Balance(balance), nil
}

func (s *storage) Withdraw(ctx context.Context, user, order string, sum int) error {
	tx, err := s.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error on begin tx: %w", err)
	}

	userID, err := s.userID(ctx, user)
	if err != nil {
		return fmt.Errorf("error on getting user id: %w", err)
	}

	orderID, err := s.orderID(ctx, order)
	if err != nil {
		return fmt.Errorf("error on getting order id: %w", err)
	}

	if _, err = tx.Exec(
		ctx,
		`INSERT INTO withdrawals (user_id, order_id, sum) VALUES ($1, $2, $3)`,
		userID,
		orderID,
		sum,
	); err != nil {
		return fmt.Errorf("error on inserting values: %w", err)
	}

	if _, err = tx.Exec(
		ctx,
		`UPDATE users SET balance = balance - $1, withdrawn = withdrawn + $1`,
		sum,
	); err != nil {
		return fmt.Errorf("error on updating values: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error on tx commit: %w", err)
	}

	return nil
}

func (s *storage) Withdrawals(ctx context.Context, user string) ([]model.Withdrawals, error) {
	withdrawals := make([]model.Withdrawals, 0)
	userID, err := s.userID(ctx, user)
	if err != nil {
		return withdrawals, fmt.Errorf("error on getting user id: %w", err)
	}

	rows, err := s.Query(
		ctx,
		`SELECT orders.number, withdrawals.sum, withdrawals.processed_at FROM withdrawals
		INNER JOIN orders ON orders.id = withdrawals.order_id
		 WHERE withdrawals.user_id = $1`,
		userID,
	)
	if err != nil {
		return withdrawals, fmt.Errorf("error on selecting values: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var w model.Withdrawals
		if err := rows.Scan(&w.Order, &w.Sum, &w.ProcessedAt); err != nil {
			return withdrawals, fmt.Errorf("error on scanning values: %w", err)
		}

		withdrawals = append(withdrawals, w)
	}

	if err := rows.Err(); err != nil {
		return withdrawals, fmt.Errorf("error on selecting values: %w", err)
	}

	return withdrawals, nil
}
