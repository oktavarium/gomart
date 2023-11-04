package orders

import "github.com/oktavarium/gomart/internal/app/internal/server/internal/model"

type Storage interface {
	GetUserByOrder(string) (string, error)
	CreateOrder(string, string, string) error
	GetOrders(string) ([]model.Order, error)
}
