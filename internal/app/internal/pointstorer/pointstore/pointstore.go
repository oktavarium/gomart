package pointstore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/oktavarium/gomart/internal/app/internal/logger"
	"github.com/oktavarium/gomart/internal/app/internal/model"
	"github.com/oktavarium/gomart/internal/app/internal/pointstorer"
)

var accrualPath = "api/orders"

type PointStore struct {
	logger      logger.Logger
	accrualAddr string
}

func NewPointStore(logger logger.Logger, addr string) *PointStore {
	return &PointStore{
		logger:      logger,
		accrualAddr: addr,
	}
}

func (ps *PointStore) GetPoints(ctx context.Context, order string) (model.Points, error) {
	var points model.Points

	client := resty.New()
	request := client.R().SetContext(ctx).SetResult(&points)
	resp, err := request.Get(fmt.Sprintf("%s/%s/%s", ps.accrualAddr, accrualPath, order))
	if err != nil {
		return points, fmt.Errorf("error on getting points from accrual system: %w", err)
	}

	switch resp.StatusCode() {
	case http.StatusNoContent:
		return points, pointstorer.ErrNotRegistered
	case http.StatusTooManyRequests:
		return points, pointstorer.ErrTooManyRequests
	case http.StatusInternalServerError:
		return points, pointstorer.ErrAccrualSystemError
	}

	return points, nil
}
