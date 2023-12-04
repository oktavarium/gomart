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

	// wait for status to update
	time.Sleep(5 * time.Second)

	testName = "get orders"
	resp, code := get(
		context.Background(),
		"orders",
		token,
		t,
	)

	require.Equal(t, http.StatusOK, code, testName)

	var orders []order
	err := json.Unmarshal(resp, &orders)
	require.NoError(t, err, testName)
	require.NotEmpty(t, orders, testName)
	require.Equal(t, goodOrderNum, orders[0].Order, testName)
	require.Equal(t, orderStatus, orders[0].Status, testName)
}
