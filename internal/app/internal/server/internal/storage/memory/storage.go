package memory

import (
	"time"

	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
)

type Storage struct {
	users Users
}

func NewStorage() *Storage {
	return &Storage{users: NewUsers()}
}

func (s *Storage) UserExists(user string) (bool, error) {
	_, ok := s.users[user]
	return ok, nil
}

func (s *Storage) RegisterUser(user, hash, salt string) error {
	s.users[user] = NewUser(hash, salt)
	return nil
}

func (s *Storage) UserHashAndSalt(user string) (string, string, error) {
	u := s.users[user]
	return u.Hash, u.Salt, nil
}

func (s *Storage) CreateOrder(user, number, status string) error {
	s.users[user].Orders[number] = Order{
		Status:     status,
		Accrual:    nil,
		UploadedAt: time.Now(),
	}

	return nil
}

func (s *Storage) GetUserByOrder(number string) (string, error) {
	for k, v := range s.users {
		if _, ok := v.Orders[number]; ok {
			return k, nil
		}
	}
	return "", nil
}

func (s *Storage) GetOrders(user string) ([]model.Order, error) {
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

func (s *Storage) GetBalance(user string) (model.Balance, error) {
	balance := s.users[user].Balance
	return model.Balance(balance), nil
}
