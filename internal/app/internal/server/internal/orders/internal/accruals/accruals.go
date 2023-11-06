package accruals

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
)

var accrualPath = "/api/orders"

type Accruals struct {
	accrualAddr string
}

func NewAccruals(accrualAddr string) *Accruals {
	return &Accruals{accrualAddr: accrualAddr}
}

func (a *Accruals) NewExecutor(orders <-chan string, bufferSize int) <-chan model.Points {
	outCh := make(chan model.Points, bufferSize)

	go func() {
		for order := range orders {
			points, err := a.getPoints(order)
			if err != nil {
				continue
			}
			outCh <- points
		}
	}()

	return outCh
}

func (a *Accruals) getPoints(order string) (model.Points, error) {
	var points model.Points

	client := resty.New()
	request := client.R().SetResult(&points)
	resp, err := request.Get(fmt.Sprintf("%s/%s", accrualPath, order))
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
