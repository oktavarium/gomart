package memory

import (
	"context"
	"fmt"
	"time"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
	"github.com/oktavarium/gomart/internal/app/internal/model"
)

type Storage struct {
	logger logger.Logger
	users  Users
}

func NewStorage(logger logger.Logger) (*Storage, error) {
	return &Storage{logger: logger, users: NewUsers()}, nil
}

func (s *Storage) UserExists(ctx context.Context, user string) (bool, error) {
	_, ok := s.users[user]
	return ok, nil
}

func (s *Storage) RegisterUser(ctx context.Context, user, hash, salt string) error {
	s.users[user] = NewUser(hash, salt)
	return nil
}

func (s *Storage) UserHashAndSalt(ctx context.Context, user string) (string, string, error) {
	u := s.users[user]
	return u.Hash, u.Salt, nil
}

func (s *Storage) NewOrder(ctx context.Context, user, number string) error {
	s.users[user].Orders[number] = Order{
		Status:     "NEW",
		Accrual:    nil,
		UploadedAt: time.Now(),
	}
	return nil
}

func (s *Storage) UpdateOrder(ctx context.Context, number, status string, accrual *int) error {
	for user, data := range s.users {
		if order, ok := data.Orders[number]; ok {
			order.Status = status
			order.Accrual = accrual

			s.users[user].Orders[number] = order
			return nil
		}
	}

	return nil
}

func (s *Storage) UserByOrder(ctx context.Context, number string) (string, error) {
	for k, v := range s.users {
		if _, ok := v.Orders[number]; ok {
			return k, nil
		}
	}
	return "", nil
}

func (s *Storage) Orders(ctx context.Context, user string) ([]model.Order, error) {
	orders := make([]model.Order, 0, len(s.users[user].Orders))
	for k, v := range s.users[user].Orders {
		order := model.Order{
			Order:      k,
			Status:     v.Status,
			Accrual:    v.Accrual,
			UploadedAt: v.UploadedAt,
		}

		orders = append(orders, order)
	}
	return orders, nil
}

func (s *Storage) OrdersByStatus(ctx context.Context, statuses []string) ([]string, error) {
	return nil, nil
}

func (s *Storage) ChekUserOrder(ctx context.Context, user, order string) error {
	if _, ok := s.users[user].Orders[order]; !ok {
		return fmt.Errorf("no such  order")
	}

	return nil
}

func (s *Storage) Balance(ctx context.Context, user string) (model.Balance, error) {
	balance := s.users[user].Balance
	return model.Balance(balance), nil
}

func (s *Storage) Withdraw(ctx context.Context, user, order string, sum int) error {
	u := s.users[user]
	u.Balance.Current -= sum
	u.Balance.Withdrawn -= sum
	u.Withdrawals = append(u.Withdrawals, Withdrawals{order, sum, time.Now()})
	s.users[user] = u

	return nil
}

func (s *Storage) Withdrawals(ctx context.Context, user string) ([]model.Withdrawals, error) {
	return nil, nil
}
