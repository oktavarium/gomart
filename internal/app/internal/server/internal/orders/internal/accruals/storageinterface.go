package accruals

import (
	"context"
)

type Storage interface {
	UpdateOrder(context.Context, string, string, *int) error
}
