package txns

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func getTransactionsFromDB() ([]Transaction, error) {
	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	if err != nil {
		return nil, err
	}

	// Create a DynamoDB client
	svc := dynamodb.New(sess)

	// Create the input for scanning the DynamoDB table
	input := &dynamodb.ScanInput{
		TableName: aws.String("Transactions"),
	}

	// Perform the scan operation
	result, err := svc.Scan(input)
	if err != nil {
		return nil, err
	}

	// Parse the scanned items into Transaction objects
	var transactions []Transaction
	for _, item := range result.Items {
		id, _ := strconv.Atoi(*item["Id"].N)
		date, _ := time.Parse("01/02/2006", *item["Date"].S)
		transaction, _ := strconv.ParseFloat(*item["Transaction"].N, 64)

		transactions = append(transactions, Transaction{
			ID:          id,
			Date:        date,
			Transaction: transaction,
		})
	}

	return transactions, nil
}
