package storage

import "github.com/oktavarium/gomart/internal/app/internal/server/shared"

type Storage interface {
	UserExists(shared.User) (bool, error)
	RegisterUser(shared.User) error
	UserInfo(shared.User) (shared.UserInfo, error)
}
