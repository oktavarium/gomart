package handlers

import (
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/auth"
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/orders"
)

type Handlers struct {
	auth   *auth.Auth
	orders *orders.Orders
}

func NewHandlers(a *auth.Auth, o *orders.Orders) *Handlers {
	return &Handlers{
		auth:   a,
		orders: o,
	}
}
