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
		errWant   error
		codeWant  int
		tokenWant bool
	}{
		{
			name:      "login invalid request",
			method:    "login",
			ct:        "application/json",
			body:      nil,
			errWant:   nil,
			codeWant:  http.StatusBadRequest,
			tokenWant: false,
		},
		{
			name:   "login invalid user",
			method: "login",
			ct:     "application/json",
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
			ct:     "application/json",
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
			ct:     "application/json",
			body: user{
				Login:    "user",
				Password: "userpass",
			},
			errWant:   nil,
			codeWant:  http.StatusOK,
			tokenWant: true,
		},
		{
			name:   "register exists user",
			method: "register",
			ct:     "application/json",
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
		_, code, token, err := post(
			context.Background(),
			test.method,
			test.ct,
			"",
			test.body,
			t,
		)
		require.Equal(t, test.errWant, err, test.name)
		require.Equal(t, test.codeWant, code, test.name)
		require.Equal(t, test.tokenWant, len(token) != 0, test.name)
	}
}
