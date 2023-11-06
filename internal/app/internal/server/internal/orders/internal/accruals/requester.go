package accruals

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
)

func getPoints(order string) (model.Points, error) {
	var points model.Points

	client := resty.New()
	request := client.R().SetDoNotParseResponse(true)
	resp, err := request.Get(fmt.Sprintf("%s/%s", accrualPath, order))
	if err != nil {
		return points, fmt.Errorf("error on getting points from accrual system: %w", err)
	}

	switch resp.StatusCode() {
	case http.StatusNoContent:
		return points, ErrNotRegistered
	case http.StatusTooManyRequests:
		body, err := io.ReadAll(resp.RawBody())
		if err != nil {
			return points, ErrReceiverError
		}
		defer resp.RawBody().Close()

		return points, ErrTooManyRequests
	case http.StatusInternalServerError:
		return points, ErrAccrualSystemError
	}

	body, err := io.ReadAll(resp.RawBody())
	if err != nil {
		return points, ErrReceiverError
	}

	defer resp.RawBody().Close()

	if err := json.Unmarshal(body, &points); err != nil {
		return points, ErrReceiverError
	}

	return points, nil
}
