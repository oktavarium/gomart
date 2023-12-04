package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func testWithdraw(t *testing.T) {
	testName := "login user for witdraw"
	_, code, token := post(
		context.Background(),
		"login",
		"application/json",
		"",
		userAndrew,
		t,
	)

	require.Equal(t, http.StatusOK, code, testName)
	require.Equal(t, len(token) != 0, true, testName)

	table := []struct {
		name     string
		method   string
		ct       string
		token    string
		body     any
		codeWant int
	}{
		{
			name:   "good withdraw",
			method: "balance/withdraw",
			ct:     "application/json",
			token:  token,
			body: withdrawal{
				Order: goodOrderNum,
				Sum:   10.0,
			},
			codeWant: http.StatusOK,
		},
		{
			name:   "withdraw unauthorized",
			method: "balance/withdraw",
			ct:     "application/json",
			token:  "",
			body: withdrawal{
				Order: goodOrderNum,
				Sum:   1.0,
			},
			codeWant: http.StatusUnauthorized,
		},
		{
			name:   "withdraw too much",
			method: "balance/withdraw",
			ct:     "application/json",
			token:  token,
			body: withdrawal{
				Order: goodOrderNum,
				Sum:   1000000.0,
			},
			codeWant: http.StatusPaymentRequired,
		},
		{
			name:   "withdraw bad order",
			method: "balance/withdraw",
			ct:     "application/json",
			token:  token,
			body: withdrawal{
				Order: badOrderNum,
				Sum:   1.0,
			},
			codeWant: http.StatusUnprocessableEntity,
		},
	}

	for _, test := range table {
		_, code, _ := post(
			context.Background(),
			test.method,
			test.ct,
			test.token,
			test.body,
			t,
		)

		require.Equal(t, test.codeWant, code, test.name)
	}
}
