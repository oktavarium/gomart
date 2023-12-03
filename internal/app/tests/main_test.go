package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
)

var endpoint string = "http://gophermart:8080/api/user"

type user struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type client struct {
	*resty.Client
	token string
}

var c *client = newClient()

func newClient() *client {
	return &client{
		Client: resty.New(),
		token:  "",
	}
}

func TestUser(t *testing.T) {
	table := []struct {
		name      string
		method    string
		body      any
		errWant   error
		codeWant  int
		tokenWant bool
	}{
		{
			name:      "login invalid request",
			method:    "login",
			body:      nil,
			errWant:   nil,
			codeWant:  http.StatusBadRequest,
			tokenWant: false,
		},
		{
			name:   "login invalid user",
			method: "login",
			body: user{
				Login:    "user",
				Password: "userpass",
			},
			errWant:   nil,
			codeWant:  http.StatusUnauthorized,
			tokenWant: false,
		},
		{
			name:   "register user",
			method: "register",
			body: user{
				Login:    "user",
				Password: "userpass",
			},
			errWant:   nil,
			codeWant:  http.StatusOK,
			tokenWant: true,
		},
		{
			name:   "login valid user",
			method: "login",
			body: user{
				Login:    "user",
				Password: "userpass",
			},
			errWant:   nil,
			codeWant:  http.StatusOK,
			tokenWant: false,
		},
		{
			name:   "register exists user",
			method: "register",
			body: user{
				Login:    "user",
				Password: "userpass",
			},
			errWant:   nil,
			codeWant:  http.StatusConflict,
			tokenWant: false,
		},
	}

	for _, test := range table {
		_, code, token, err := c.post(
			context.Background(),
			test.method,
			nil,
			t,
		)
		require.Equal(t, test.errWant, err, test.name)
		require.Equal(t, test.codeWant, code, test.name)
		require.Equal(t, test.tokenWant, len(token) != 0, test.name)
	}
}

func TestOrders(t *testing.T) {
	_, code, token, err := c.post(
		context.Background(),
		"register",
		user{
			Login:    "andrew",
			Password: "userpass",
		},
		t,
	)

	require.Equal(t, nil, err, "registering user for orders")
	require.Equal(t, http.StatusOK, code, "registering user for orders")
	require.Equal(t, len(token) != 0, true, "registering user for orders")

	table := struct {
		name     string
		method   string
		order    string
		errWant  error
		codeWant int
		respWant string
	} {
		{
			name: "create wrong order",
			method: "orders",
			order: "0",
			errWant: nil,
			codeWant: http.Status
		},
	}
}

func (c *client) post(
	ctx context.Context,
	method string,
	body any,
	t *testing.T,
) (any, int, string, error) {

	t.Helper()

	var result interface{}
	request := c.R().SetContext(ctx).SetBody(body).SetResult(&result)
	resp, err := request.Post(fmt.Sprintf("%s/%s", endpoint, method))
	if err != nil {
		return result, 0, "", fmt.Errorf("error on making post request: %w", err)
	}

	token := resp.Header().Get("Authorization")

	return result, resp.StatusCode(), token, nil
}
