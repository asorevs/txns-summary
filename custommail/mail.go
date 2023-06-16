package custommail

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"strconv"

	"github.com/asorevs/txnsummary/config"
	"github.com/asorevs/txnsummary/txns"
	"gopkg.in/gomail.v2"
)

func SendCustomMail() {
	// Retrieve configuration settings
	cfg := config.NewConfig()

	// Create a new email message.
	m := gomail.NewMessage()

	m.SetHeader("From", cfg.EmailSender)
	m.SetHeader("To", cfg.EmailRecipient)
	m.SetHeader("Subject", "Transaction summary")

	// Parse the HTML template.
	t, err := template.ParseFiles("custommail/template.html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Create a buffer to hold the rendered template.
	var body bytes.Buffer

	// Attach the image to the email and get the Content-ID.
	imagePath := "static/images/stori.jpeg"
	imageID := filepath.Base(imagePath)
	m.Embed(imagePath)

	// Analyze transactions and retrieve the data
	totalBalance, averageCreditAmount, averageDebitAmount, transactionData := txns.AnalyseTxns()

	// Execute the template with the provided data and write the output to the buffer.
	t.Execute(&body, struct {
		Name                string
		Message             string
		ImageID             string
		TotalBalance        string
		AverageCreditAmount string
		AverageDebitAmount  string
		TransactionData     []txns.MonthTxns
	}{
		Name:                "Alberto",
		Message:             "Thank you for trusting us to help you achieve your goals. Here's a brief summary of your transactions over the past months. If you have any further questions or need assistance, please don't hesitate to reach out. We're here to help. ",
		ImageID:             imageID,
		TotalBalance:        strconv.FormatFloat(totalBalance, 'f', 2, 64),
		AverageCreditAmount: strconv.FormatFloat(averageCreditAmount, 'f', 2, 64),
		AverageDebitAmount:  strconv.FormatFloat(averageDebitAmount, 'f', 2, 64),
		TransactionData:     transactionData,
	})

	// Set the body of the email.
	m.SetBody("text/html", body.String())

	// Create a new SMTP dialer.
	d := gomail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.EmailSender, cfg.EmailPassword)

	// Send the email using the SMTP dialer.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nEmail Sent!")
}
