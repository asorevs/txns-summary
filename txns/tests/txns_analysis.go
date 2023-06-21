package test

import (
	"testing"

	"github.com/asorevs/txnsummary/txns"
)

func TestCalculateTotalBalance(t *testing.T) {
	// Create sample transactions for testing
	transactions := []txns.Transaction{
		{Transaction: 100},
		{Transaction: -50},
		{Transaction: 75},
	}

	totalBalance := txns.CalculateTotalBalance(transactions)

	expectedTotalBalance := 125.0
	if totalBalance != expectedTotalBalance {
		t.Errorf("Expected total balance %.2f, got %.2f", expectedTotalBalance, totalBalance)
	}
}
