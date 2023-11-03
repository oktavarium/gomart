package storage

import "github.com/oktavarium/gomart/internal/app/internal/server/internal/model"

type Storage interface {
	UserExists(model.User) (bool, error)
	RegisterUser(model.User) error
	UserInfo(model.User) (model.UserInfo, error)
}
