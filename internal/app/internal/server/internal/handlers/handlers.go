package handlers

import (
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/auth"
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/orders"
)

type Handlers struct {
	auth        *auth.Auth
	orders      *orders.Orders
	accrualAddr string
}

func NewHandlers(a *auth.Auth, o *orders.Orders, accuralAddr string) *Handlers {
	return &Handlers{a, o, accuralAddr}
}
