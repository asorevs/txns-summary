package txns

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Count of transactions for a specific month.
type MonthTxns struct {
	Month string
	Count int
}

// Summarizes transactions and returns: totalBalance, averageCreditAmount, averageDebitAmount and transactionData by month
func AnalyseTxns() (float64, float64, float64, []MonthTxns) {
	// Read the transactions from the CSV file
	transactions, err := readTxnsFromFile("txns.csv")
	if err != nil {
		fmt.Println("Error reading transactions:", err)
		return 0, 0, 0, nil
	}

	// Calculate the total balance
	totalBalance := calculateTotalBalance(transactions)
	fmt.Printf("Total balance is: %.2f\n", totalBalance)

	// Calculate the average debit and credit amounts
	averageDebitAmount := calculateAverageAmount(transactions, "debit")
	averageCreditAmount := calculateAverageAmount(transactions, "credit")
	fmt.Printf("Average debit amount: %.2f\n", averageDebitAmount)
	fmt.Printf("Average credit amount: %.2f\n", averageCreditAmount)

	// Calculate the number of transactions for each month
	fmt.Printf("\nNumber of transactions\n")
	transactionsByMonth := countTxnsByMonth(transactions)
	for _, mc := range transactionsByMonth {
		monthName := mc.Month
		count := mc.Count
		fmt.Printf("%s: %d\n", monthName, count)
	}

	return totalBalance, averageCreditAmount, averageDebitAmount, transactionsByMonth
}

// Reads the transactions from the specified CSV file.
func readTxnsFromFile(filename string) ([]Transaction, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var transactions []Transaction
	for _, record := range records[1:] {
		id, _ := strconv.Atoi(record[0])
		date, _ := time.Parse("01/02/2006", record[1])
		transaction, _ := strconv.ParseFloat(record[2], 64)

		transactions = append(transactions, Transaction{
			ID:          id,
			Date:        date,
			Transaction: transaction,
		})
	}

	return transactions, nil
}

// Calculates the total balance from the given transactions.
func calculateTotalBalance(transactions []Transaction) float64 {
	totalBalance := 0.0
	for _, txn := range transactions {
		totalBalance += txn.Transaction
	}
	return totalBalance
}

// Counts the number of transactions for each month and returns them sorted by date.
func countTxnsByMonth(transactions []Transaction) []MonthTxns {
	transactionsByMonth := make(map[time.Time]int)
	months := make([]time.Time, 0)

	for _, txn := range transactions {
		// Truncate the date to the month level
		currentMonth := time.Date(txn.Date.Year(), txn.Date.Month(), 1, 0, 0, 0, 0, txn.Date.Location())
		transactionsByMonth[currentMonth]++

		// Collects distinct months to ensure that results are later displayed in chronological order
		found := false
		for _, distinctMonth := range months {
			if distinctMonth.Equal(currentMonth) {
				found = true
				break
			}
		}
		if !found {
			months = append(months, currentMonth)
		}
	}

	// Sort the unique months slice in chronological order
	sort.Slice(months, func(i, j int) bool {
		return months[i].Before(months[j])
	})

	// Create a slice of MonthTxns structs
	monthTxnsSlice := make([]MonthTxns, 0, len(months))
	for _, month := range months {
		monthStr := month.Format("January")
		count := transactionsByMonth[month]
		monthTxnsSlice = append(monthTxnsSlice, MonthTxns{
			Month: monthStr,
			Count: count,
		})
	}

	return monthTxnsSlice
}

// Calculates the average amount for the specified transaction type.
func calculateAverageAmount(transactions []Transaction, txnType string) float64 {
	count := 0
	totalAmount := 0.0
	for _, txn := range transactions {
		if strings.ToLower(txnType) == "debit" && txn.Transaction < 0 {
			count++
			totalAmount += txn.Transaction
		} else if strings.ToLower(txnType) == "credit" && txn.Transaction > 0 {
			count++
			totalAmount += txn.Transaction
		}
	}

	if count > 0 {
		return totalAmount / float64(count)
	}
	return 0.0
}
