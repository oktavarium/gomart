package pointstorer

import (
	"context"

	"github.com/oktavarium/gomart/internal/app/internal/model"
)

type PointStorer interface {
	GetPoints(context.Context, string) (model.Points, error)
}
