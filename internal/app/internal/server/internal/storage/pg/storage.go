package pg

import (
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
)

type storage struct {
}

func NewStorage(dbURI string) (*storage, error) {
	s := &storage{}

	return s, nil
}

func (s *storage) UserExists(u model.User) (bool, error) {
	return false, nil
}

func (s *storage) RegisterUser(u model.User) error {
	return nil
}

func (s *storage) UserInfo(u model.User) (model.UserInfo, error) {
	return model.UserInfo{}, nil
}
