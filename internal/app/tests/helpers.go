package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
)

var endpoint string = "http://gophermart:8080/api/user"

type user struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type order struct {
	Order      string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    float32   `json:"accrual"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type balance struct {
	Current   float32 `json:"current"`
	Withdrawn float32 `json:"withdrawn"`
}

type withdrawal struct {
	Order string  `json:"order"`
	Sum   float32 `json:"sum"`
}

type withdrawals struct {
	Order       string    `json:"order"`
	Sum         float32   `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}

func post(
	ctx context.Context,
	method string,
	ct string,
	token string,
	body any,
	t *testing.T,
) (any, int, string, error) {

	t.Helper()

	var result interface{}
	c := resty.New()
	request := c.R().SetContext(ctx).SetBody(body).SetResult(&result).SetHeader("Content-Type", ct)
	if len(token) != 0 {
		request.SetHeader("Authorization", token)
	}

	resp, err := request.Post(fmt.Sprintf("%s/%s", endpoint, method))
	if err != nil {
		return result, 0, "", fmt.Errorf("error on making post request: %w", err)
	}

	return result, resp.StatusCode(), resp.Header().Get("Authorization"), nil
}

func get(
	ctx context.Context,
	method string,
	token string,
	t *testing.T,
) ([]byte, int, error) {

	t.Helper()

	c := resty.New()
	request := c.R().SetContext(ctx)
	if len(token) != 0 {
		request.SetHeader("Authorization", token)
	}

	resp, err := request.Get(fmt.Sprintf("%s/%s", endpoint, method))
	if err != nil {
		return nil, 0, fmt.Errorf("error on making get request: %w", err)
	}

	return resp.Body(), resp.StatusCode(), nil
}
