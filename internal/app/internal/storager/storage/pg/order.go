package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/oktavarium/gomart/internal/app/internal/model"
)

func (s *storage) NewOrder(ctx context.Context, user, number string) error {
	tx, err := s.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error on begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	userID, err := s.userID(ctx, user)
	if err != nil {
		return fmt.Errorf("error on getting user id: %w", err)
	}

	_, err = tx.Exec(ctx, `INSERT INTO orders (user_id, number) VALUES ($1, $2)`, userID, number)
	if err != nil {
		return fmt.Errorf("error on inserting values: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error on tx commit: %w", err)
	}

	return nil
}

func (s *storage) UpdateOrder(ctx context.Context, number, status string, accrual float32) error {
	tx, err := s.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error on begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	user, err := s.UserByOrder(ctx, number)
	if err != nil {
		return fmt.Errorf("error on getting user by order: %w", err)
	}

	_, err = tx.Exec(ctx, `UPDATE users SET balance = balance + $1 WHERE name = $2`, accrual, user)
	if err != nil {
		return fmt.Errorf("error on updating user balance %w", err)
	}

	_, err = tx.Exec(
		ctx,
		`UPDATE orders SET status = $1, accrual = $2 WHERE number = $3`,
		status,
		accrual,
		number,
	)

	if err != nil {
		return fmt.Errorf("error on updating values: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error on tx commit: %w", err)
	}

	return nil
}

func (s *storage) Orders(ctx context.Context, user string) ([]model.Order, error) {
	orders := make([]model.Order, 0)
	userID, err := s.userID(ctx, user)
	if err != nil {
		return orders, fmt.Errorf("error on getting userID: %w", err)
	}

	rows, err := s.Query(
		ctx,
		`SELECT number, status, accrual, uploaded_at FROM orders WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return orders, fmt.Errorf("error on selecting values: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order model.Order
		err := rows.Scan(&order.Order, &order.Status, &order.Accrual, &order.UploadedAt)
		if err != nil {
			return orders, fmt.Errorf("error on scanning values: %w", err)
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		if err == pgx.ErrNoRows {
			return orders, nil
		}

		return orders, fmt.Errorf("error on selectins values: %w", err)
	}

	return orders, nil
}

func (s *storage) OrdersByStatus(ctx context.Context, statuses []string) ([]string, error) {
	orders := make([]string, 0)
	statusStr := strings.Join(statuses, ",")

	rows, err := s.Query(ctx, `SELECT number FROM orders WHERE status IN ($1)`, statusStr)
	if err != nil {
		return orders, fmt.Errorf("error on selecting values: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order string
		if err := rows.Scan(&order); err != nil {
			return orders, fmt.Errorf("error on scanning values: %w", err)
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return orders, fmt.Errorf("error on selectins values: %w", err)
	}

	return orders, nil
}

func (s *storage) ChekUserOrder(ctx context.Context, user, order string) error {
	userID, err := s.userID(ctx, user)
	if err != nil {
		return fmt.Errorf("error on getting user id: %w", err)
	}

	var id string
	row := s.QueryRow(ctx, `SELECT id FROM orders WHERE user_id = $1, number = $2`, userID, order)
	err = row.Scan(&id)
	if err != nil {
		if err != pgx.ErrNoRows {
			return fmt.Errorf("error on selecting values: %w", err)
		}
	}

	return nil
}
