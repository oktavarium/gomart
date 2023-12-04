package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func testWithdrawals(t *testing.T) {
	testName := "login user for withdrawals"
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
		token    string
		bodyWant bool
		codeWant int
	}{
		{
			name:     "get withdrawals",
			method:   "withdrawals",
			token:    tokenAndrew,
			bodyWant: true,
			codeWant: http.StatusOK,
		},
		{
			name:     "get empty withdrawals",
			method:   "withdrawals",
			token:    tokenJimmy,
			bodyWant: false,
			codeWant: http.StatusNoContent,
		},
		{
			name:     "withdrawals unauthorized",
			method:   "withdrawals",
			token:    "",
			bodyWant: false,
			codeWant: http.StatusUnauthorized,
		},
	}

	for _, test := range table {
		resp, code := get(
			context.Background(),
			test.method,
			test.token,
			t,
		)

		require.Equal(t, test.codeWant, code, test.name)

		if test.bodyWant {
			var w []withdrawals
			err := json.Unmarshal(resp, &w)
			require.NoError(t, err, test.name)
			require.NotEmpty(t, w, test.name)
		}
	}
}
