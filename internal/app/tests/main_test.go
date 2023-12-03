package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/oktavarium/gomart/internal/app/internal/pointstorer"
	"github.com/stretchr/testify/require"
)

type client struct {
	*resty.Client
}

var c *client

func newClient() *client {
	return &client{resty.New()}
}

func TestMain(t *testing.T) {
	c = newClient()
	TestInvalidUser(t)

}

func TestInvalidUser(t *testing.T) {
	res, err := c.post("http://127.0.0.1:8080/api/user/register", context.Background(), t)
	fmt.Println(res)
	require.Equal(t, err, nil)
}

func (c *client) post(endpoint string, ctx context.Context, t *testing.T) (any, error) {
	t.Helper()

	var result interface{}
	request := c.R().SetContext(ctx).SetResult(&result)
	resp, err := request.Post(endpoint)
	if err != nil {
		return result, fmt.Errorf("error on getting points from accrual system: %w", err)
	}

	switch resp.StatusCode() {
	case http.StatusNoContent:
		return result, fmt.Errorf("error on getting points: %w", pointstorer.ErrNotRegistered)

	case http.StatusTooManyRequests:
		return result, fmt.Errorf(
			"%w, response body: %s",
			pointstorer.ErrTooManyRequests,
			string(resp.Body()),
		)
	case http.StatusInternalServerError:
		return result, fmt.Errorf("error on getting points: %w", pointstorer.ErrAccrualSystemError)
	}

	return result, nil
}
