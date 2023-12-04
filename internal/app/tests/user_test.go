package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func testUser(t *testing.T) {
	table := []struct {
		name      string
		method    string
		ct        string
		body      any
		codeWant  int
		tokenWant bool
	}{
		{
			name:      "login invalid request",
			method:    "login",
			ct:        "application/json",
			body:      nil,
			codeWant:  http.StatusBadRequest,
			tokenWant: false,
		},
		{
			name:      "login invalid user",
			method:    "login",
			ct:        "application/json",
			body:      userAndrew,
			codeWant:  http.StatusUnauthorized,
			tokenWant: false,
		},
		{
			name:      "register user",
			method:    "register",
			ct:        "application/json",
			body:      userAndrew,
			codeWant:  http.StatusOK,
			tokenWant: true,
		},
		{
			name:      "login valid user",
			method:    "login",
			ct:        "application/json",
			body:      userAndrew,
			codeWant:  http.StatusOK,
			tokenWant: true,
		},
		{
			name:      "login valid user",
			method:    "login",
			ct:        "application/json",
			body:      userJimmy,
			codeWant:  http.StatusOK,
			tokenWant: true,
		},
		{
			name:      "register exists user",
			method:    "register",
			ct:        "application/json",
			body:      userAndrew,
			codeWant:  http.StatusConflict,
			tokenWant: false,
		},
		{
			name:      "register bad request",
			method:    "register",
			ct:        "",
			body:      nil,
			codeWant:  http.StatusBadRequest,
			tokenWant: false,
		},
	}

	for _, test := range table {
		_, code, token := post(
			context.Background(),
			test.method,
			test.ct,
			"",
			test.body,
			t,
		)
		require.Equal(t, test.codeWant, code, test.name)
		require.Equal(t, test.tokenWant, len(token) != 0, test.name)
	}
}
