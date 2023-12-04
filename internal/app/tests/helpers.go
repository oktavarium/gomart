package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
)

var endpoint string = "http://gophermart:8080/api/user"

var userAndrew = user{
	Login:    "andrew",
	Password: "userpass",
}

var userJimmy = user{
	Login:    "jimmy",
	Password: "userpass",
}

var goodOrderNum string = "12345678903"
var badOrderNum string = "12345678904"
var orderStatus string = "PROCESSED"

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
	m string,
	ct string,
	token string,
	b any,
	t *testing.T,
) (any, int, string) {
	t.Helper()

	var result interface{}
	c := resty.New()
	request := c.R().SetContext(ctx).SetBody(b).SetResult(&result).SetHeader("Content-Type", ct)
	if len(token) != 0 {
		request.SetHeader("Authorization", token)
	}

	resp, err := request.Post(fmt.Sprintf("%s/%s", endpoint, m))
	require.NoError(t, err, "making get request")

	return result, resp.StatusCode(), resp.Header().Get("Authorization")
}

func get(ctx context.Context, m string, token string, t *testing.T) ([]byte, int) {
	t.Helper()

	c := resty.New()
	request := c.R().SetContext(ctx)
	if len(token) != 0 {
		request.SetHeader("Authorization", token)
	}

	resp, err := request.Get(fmt.Sprintf("%s/%s", endpoint, m))
	require.NoError(t, err, "making get request")

	return resp.Body(), resp.StatusCode()
}
