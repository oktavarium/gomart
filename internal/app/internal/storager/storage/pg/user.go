package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (s *storage) UserExists(ctx context.Context, user string) (bool, error) {
	var id string
	row := s.QueryRow(ctx, `SELECT id FROM users WHERE name = $1`, user)
	err := row.Scan(&id)
	if err != nil {
		if err != pgx.ErrNoRows {
			return false, fmt.Errorf("error on selecting values: %w", err)
		} else {
			return false, nil
		}
	}

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
		`INSERT INTO users (name, hash, salt) VALUES ($1, $2, $3)`,
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

func (s *storage) GetUserHashAndSalt(ctx context.Context, user string) (string, string, error) {
	var hash, salt string
	row := s.QueryRow(
		ctx,
		`SELECT hash, salt FROM users WHERE name = $1`,
		user,
	)

	if err := row.Scan(&hash, &salt); err != nil {
		return hash, salt, fmt.Errorf("error on scanning values: %w", err)
	}

	return hash, salt, nil
}

func (s *storage) GetUserByOrder(ctx context.Context, number string) (string, error) {
	var user string
	row := s.QueryRow(
		ctx,
		`SELECT users.name FROM orders
	INNER JOIN users ON users.id = orders.user_id
	WHERE orders.number = $1`,
		number,
	)

	if err := row.Scan(&user); err != nil {
		if err != pgx.ErrNoRows {
			return user, fmt.Errorf("error on scanning values: %w", err)
		} else {
			return user, nil
		}

	}

	return user, nil
}

func (s *storage) getUserID(ctx context.Context, user string) (string, error) {
	var userID string
	row := s.QueryRow(ctx, `SELECT id FROM users WHERE name = $1`, user)

	if err := row.Scan(&userID); err != nil {
		return userID, fmt.Errorf("error on scanning values: %w", err)
	}

	return userID, nil
}
