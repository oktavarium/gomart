package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func testGetOrders(t *testing.T) {
	testName := "login user for get orders"
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

	// wait for status to update
	time.Sleep(5 * time.Second)

	testName = "get orders"
	resp, code, err := get(
		context.Background(),
		"orders",
		token,
		t,
	)

	require.Equal(t, nil, err, testName)
	require.Equal(t, http.StatusOK, code, testName)

	var orders []order
	err = json.Unmarshal(resp, &orders)
	require.NoError(t, err, testName)
	require.NotEmpty(t, orders, testName)
	require.Equal(t, "12345678903", orders[0].Order, testName)
	require.Equal(t, "PROCESSED", orders[0].Status, testName)
}
