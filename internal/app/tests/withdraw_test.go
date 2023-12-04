package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func testWithdraw(t *testing.T) {
	testName := "login user for witdraw"
	_, code, token, err := post(
		context.Background(),
		"login",
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

	testName = "withdraw"
	_, code, _, err = post(
		context.Background(),
		"balance/withdraw",
		"application/json",
		token,
		withdrawal{
			Order: "12345678903",
			Sum:   10.0,
		},
		t,
	)

	require.Equal(t, nil, err, testName)
	require.Equal(t, http.StatusOK, code, testName)

	require.NoError(t, err, testName)

	resp, code, err := get(
		context.Background(),
		"withdrawals",
		token,
		t,
	)

	require.Equal(t, nil, err, testName)
	require.Equal(t, http.StatusOK, code, testName)

	var w withdrawals
	err = json.Unmarshal(resp, &w)
	require.NoError(t, err, testName)
	require.NotEmpty(t, w, testName)
}
