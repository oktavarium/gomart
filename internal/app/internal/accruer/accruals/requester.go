package accruals

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/oktavarium/gomart/internal/app/internal/model"
)

func (a *Accruals) getPoints(order string) (model.Points, error) {
	var points model.Points
	a.logger.Info("GET POINTS")
	client := resty.New()
	request := client.R().SetResult(&points)
	resp, err := request.Get(fmt.Sprintf("%s/%s/%s", a.accrualAddr, accrualPath, order))
	if err != nil {
		return points, fmt.Errorf("error on getting points from accrual system: %w", err)
	}
	a.logger.Info(fmt.Sprintf("%s %d", "REQUEST MADE", resp.StatusCode()))
	switch resp.StatusCode() {
	case http.StatusNoContent:
		return points, ErrNotRegistered
	case http.StatusTooManyRequests:
		body, err := io.ReadAll(resp.RawBody())
		if err != nil {
			return points, ErrReceiverError
		}
		fmt.Println("BODY!!!", body)
		defer resp.RawBody().Close()

		return points, ErrTooManyRequests
	case http.StatusInternalServerError:
		return points, ErrAccrualSystemError
	}

	// body, err := io.ReadAll(resp.RawBody())
	// if err != nil {
	// 	return points, ErrReceiverError
	// }

	// defer resp.RawBody().Close()

	// if err := json.Unmarshal(body, &points); err != nil {
	// 	return points, ErrReceiverError
	// }

	return points, nil
}
