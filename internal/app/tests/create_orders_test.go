package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func testCreateOrders(t *testing.T) {
	testName := "register user for orders"
	_, code, token, err := post(
		context.Background(),
		"register",
		"application/json",
		"",
		user{
			Login:    "andrew",
			Password: "userpass",
		},
		t,
	)

	require.Equal(t, nil, err, testName)
	require.Equal(t, http.StatusOK, code, testName)
	require.Equal(t, len(token) != 0, true, testName)

	table := []struct {
		name     string
		method   string
		ct       string
		order    string
		token    string
		errWant  error
		codeWant int
		respWant any
	}{
		{
			name:     "create wrong order",
			method:   "orders",
			ct:       "text/plain",
			order:    "12345678904",
			token:    token,
			errWant:  nil,
			codeWant: http.StatusUnprocessableEntity,
			respWant: nil,
		},
		{
			name:     "create good order unauthorized",
			method:   "orders",
			ct:       "text/plain",
			order:    "12345678903",
			token:    "",
			errWant:  nil,
			codeWant: http.StatusUnauthorized,
			respWant: nil,
		},
		{
			name:     "create good order",
			method:   "orders",
			ct:       "text/plain",
			order:    "12345678903",
			token:    token,
			errWant:  nil,
			codeWant: http.StatusAccepted,
			respWant: nil,
		},
	}

	for _, test := range table {
		resp, code, _, err := post(
			context.Background(),
			test.method,
			test.ct,
			test.token,
			test.order,
			t,
		)
		require.Equal(t, test.errWant, err, test.name)
		require.Equal(t, test.codeWant, code, test.name)
		require.Equal(t, test.respWant, resp, test.name)
	}
}
