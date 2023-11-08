package pg

import (
	"context"
	"fmt"
)

func (s *storage) UserExists(ctx context.Context, user string) (bool, error) {
	return true, nil
}

func (s *storage) RegisterUser(ctx context.Context, user, hash, salt string) error {
	tx, err := s.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error on begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(
		ctx,
		`INSERT INTO users (user, hash, salt) VALUES ($1, $1, $3)`,
		user,
		hash,
		salt,
	); err != nil {
		return fmt.Errorf("error on making insert: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error on tx commit: %w", err)
	}

	return nil
}

func (s *storage) UserHashAndSalt(ctx context.Context, user string) (string, string, error) {
	var hash, salt string
	row := s.QueryRow(
		ctx,
		`SELECT hash, salt FROM users WHERE user = $1`,
		user,
	)

	if err := row.Scan(&hash, &salt); err != nil {
		return hash, salt, fmt.Errorf("error on scanning values: %w", err)
	}

	return hash, salt, nil
}

func (s *storage) UserByOrder(ctx context.Context, number string) (string, error) {
	var user string
	row := s.QueryRow(
		ctx,
		`SELECT users.user FROM orders
	INNER JOIN users ON users.id = orders.user_id
	WHERE orders.number = $1`,
		number,
	)

	if err := row.Scan(&user); err != nil {
		return user, fmt.Errorf("error on scanning values: %w", err)
	}

	return user, nil
}

func (s *storage) userId(ctx context.Context, user string) (string, error) {
	var userId string
	row := s.QueryRow(ctx, `SELECT id FROM users WHERE user = $1`, user)

	if err := row.Scan(&userId); err != nil {
		return userId, fmt.Errorf("error on scanning values: %w", err)
	}

	return userId, nil
}
