package memory

import (
	"fmt"

	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
)

type storage struct {
	users map[string]model.User
}

func NewStorage() *storage {
	return &storage{users: make(map[string]model.User)}
}

func (s *storage) UserExists(u model.User) (bool, error) {
	_, ok := s.users[u.Login]
	return ok, nil
}

func (s *storage) RegisterUser(u model.User) error {
	exists, err := s.UserExists(u)
	if err != nil {
		return fmt.Errorf("error on checking user existance: %w", err)
	}

	if exists {
		return fmt.Errorf("user already exists: %w", err)
	}

	s.users[u.Login] = u
	return nil
}

func (s *storage) UserInfo(u model.User) (model.UserInfo, error) {
	exists, err := s.UserExists(u)
	if err != nil {
		return model.UserInfo{}, fmt.Errorf("error on checking user existance: %w", err)
	}

	if !exists {
		return model.UserInfo{}, fmt.Errorf("user not exists: %w", err)
	}

	return s.users[u.Login].Info, nil
}
