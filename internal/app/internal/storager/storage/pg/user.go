package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (s *storage) UserExists(user string) bool {
	if _, err := s.getUserID(user); err != nil {
		return false
	}

	return true
}

func (s *storage) RegisterUser(ctx context.Context, user, hash, salt string) error {
	tx, err := s.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error on begin tx: %w", err)
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			s.logger.Error(err)
		}
	}()

	var id string
	if err := tx.QueryRow(
		ctx,
		`INSERT INTO users (name, hash, salt) VALUES ($1, $2, $3) RETURNING id`,
		user,
		hash,
		salt,
	).Scan(&id); err != nil {
		return fmt.Errorf("error on making insert: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error on tx commit: %w", err)
	}

	s.users[user] = id

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

func (s *storage) getUserID(user string) (string, error) {
	if userID, ok := s.users[user]; ok {
		return userID, nil
	}

	return "", fmt.Errorf("no such user")
}

func (s *storage) cacheUsers(ctx context.Context) error {
	rows, err := s.Query(ctx, `SELECT id, name FROM users`)
	if err != nil {
		return fmt.Errorf("error on selecting users: %w", err)
	}

	defer rows.Close()

	var name, id string
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			return fmt.Errorf("error on scanning values: %w", err)
		}

		s.users[name] = id
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error on selecting values: %w", err)
	}

	return nil
}
