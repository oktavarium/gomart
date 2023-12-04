package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func testCreateOrders(t *testing.T) {
	testName := "login user for orders"
	_, code, tokenAndrew := post(
		context.Background(),
		"login",
		"application/json",
		"",
		userAndrew,
		t,
	)

	require.Equal(t, http.StatusOK, code, testName)
	require.Equal(t, len(tokenAndrew) != 0, true, testName)

	_, code, tokenJimmy := post(
		context.Background(),
		"login",
		"application/json",
		"",
		userJimmy,
		t,
	)

	require.Equal(t, http.StatusOK, code, testName)
	require.Equal(t, len(tokenJimmy) != 0, true, testName)

	table := []struct {
		name     string
		method   string
		ct       string
		order    string
		token    string
		codeWant int
		respWant any
	}{
		{
			name:     "create wrong order",
			method:   "orders",
			ct:       "text/plain",
			order:    badOrderNum,
			token:    tokenAndrew,
			codeWant: http.StatusUnprocessableEntity,
		},
		{
			name:     "create order bad request",
			method:   "orders",
			ct:       "",
			order:    badOrderNum,
			token:    tokenAndrew,
			codeWant: http.StatusBadRequest,
		},
		{
			name:     "create good order unauthorized",
			method:   "orders",
			ct:       "text/plain",
			order:    goodOrderNum,
			token:    "",
			codeWant: http.StatusUnauthorized,
		},
		{
			name:     "create good order",
			method:   "orders",
			ct:       "text/plain",
			order:    goodOrderNum,
			token:    tokenAndrew,
			codeWant: http.StatusAccepted,
		},
		{
			name:     "create good order with same number",
			method:   "orders",
			ct:       "text/plain",
			order:    goodOrderNum,
			token:    tokenAndrew,
			codeWant: http.StatusOK,
		},
		{
			name:     "create good order with same number from another user",
			method:   "orders",
			ct:       "text/plain",
			order:    goodOrderNum,
			token:    tokenJimmy,
			codeWant: http.StatusConflict,
		},
	}

	for _, test := range table {
		_, code, _ := post(context.Background(), test.method, test.ct, test.token, test.order, t)
		require.Equal(t, test.codeWant, code, test.name)
	}
}
