package pointmock

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/oktavarium/gomart/internal/app/internal/logger"
	"github.com/oktavarium/gomart/internal/app/internal/model"
)

type pointmock struct {
	logger logger.Logger
}

func NewPointmock(logger logger.Logger) *pointmock {
	return &pointmock{logger: logger}
}

func (pm *pointmock) GetPoints(ctx context.Context, order string) (model.Points, error) {
	pm.logger.Debug(fmt.Sprintf("getting points for order: %s", order))
	var p model.Points

	p.Accrual = 10 + rand.Float32()
	p.Order = order
	p.Status = "PROCESSED"

	return p, nil
}
