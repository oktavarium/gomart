package accruals

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/oktavarium/gomart/internal/app/internal/model"
)

func (a *Accruals) getPoints(order string) (model.Points, error) {
	var points model.Points

	client := resty.New()
	request := client.R().SetResult(&points)
	resp, err := request.Get(fmt.Sprintf("%s/%s/%s", a.accrualAddr, accrualPath, order))
	if err != nil {
		return points, fmt.Errorf("error on getting points from accrual system: %w", err)
	}

	switch resp.StatusCode() {
	case http.StatusNoContent:
		return points, ErrNotRegistered
	case http.StatusTooManyRequests:
		return points, ErrTooManyRequests
	case http.StatusInternalServerError:
		return points, ErrAccrualSystemError
	}

	return points, nil
}
