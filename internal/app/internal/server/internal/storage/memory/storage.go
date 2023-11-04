package memory

import (
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
)

type Storage struct {
	users map[string]model.User
}

func NewStorage() *Storage {
	return &Storage{users: make(map[string]model.User)}
}

func (s *Storage) UserExists(user string) (bool, error) {
	return true, nil
}

func (s *Storage) RegisterUser(user, hash, salt string) error {
	return nil
}

func (s *Storage) UserHashAndSalt(user string) (string, string, error) {
	return "", "", nil
}

func (s *Storage) NewOrder(user, order string) error {
	return nil
}
