package tests

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	if os.Getenv("TEST_MODE") == "" {
		t.Skip("Skipping testing in CI environment")
	}

	testUser(t)
	testCreateOrders(t)
	testGetOrders(t)
	testBalance(t)
	testWithdraw(t)
	testWithdrawals(t)
}
