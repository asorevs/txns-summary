package txns

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GenerateRandomTxns() {
	rand.Seed(time.Now().UnixNano())

	// Create a new csv file
	file, err := os.Create("txns.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	header := "Id,Date,Transaction\n"
	_, err = file.WriteString(header)
	if err != nil {
		fmt.Println("Error writing header:", err)
		return
	}

	// Generate random data for the last few months
	now := time.Now()
	monthsAgo := now.AddDate(0, -5, 0)
	daysInRange := int(now.Sub(monthsAgo).Hours() / 24)
	id := 0
	for i := 0; i < daysInRange; i++ {
		numTransactions := rand.Intn(10)

		// Generate transactions for the current day
		date := now.AddDate(0, 0, -i)
		for j := 0; j < numTransactions; j++ {
			transaction := randomFloat(-100, 100, 2)
			record := Transaction{
				ID:          id,
				Date:        date,
				Transaction: transaction,
			}
			id++

			recordStr := fmt.Sprintf("%d,%s,%.2f\n", record.ID, record.Date.Format("01/02/2006"), record.Transaction)
			_, err = file.WriteString(recordStr)
			if err != nil {
				fmt.Println("Error writing record:", err)
				return
			}
		}
	}

	fmt.Println("Random data generated and saved to txns.csv")
}

// Generates a random float64 between min and max with the specified number of decimal places.
func randomFloat(min, max float64, decimalPlaces int) float64 {
	value := min + rand.Float64()*(max-min)
	rounded := strconv.FormatFloat(value, 'f', decimalPlaces, 64)
	result, _ := strconv.ParseFloat(rounded, 64)
	return result
}
