package pg

import (
	"github.com/oktavarium/gomart/internal/app/internal/server/shared"
)

type storage struct {
}

func NewStorage(dbURI string) (*storage, error) {
	s := &storage{}

	return s, nil
}

func (s *storage) UserExists(u shared.User) (bool, error) {
	return false, nil
}

func (s *storage) RegisterUser(u shared.User) error {
	return nil
}

func (s *storage) UserInfo(u shared.User) (shared.UserInfo, error) {
	return shared.UserInfo{}, nil
}
