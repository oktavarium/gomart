package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func testBalance(t *testing.T) {
	testName := "login user for get balance"
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

	testName = "get balance"
	resp, code := get(
		context.Background(),
		"balance",
		token,
		t,
	)

	require.Equal(t, http.StatusOK, code, testName)

	var b balance
	err := json.Unmarshal(resp, &b)
	require.NoError(t, err, testName)
	require.NotEmpty(t, b.Current, testName)
	require.Empty(t, b.Withdrawn, testName)
}
