package config

import (
	"os"
	"strconv"
)

// Represents the application configuration.
type Config struct {
	EmailSender    string
	EmailPassword  string
	EmailRecipient string
	SMTPHost       string
	SMTPPort       int
}

// Creates a new Config instance with values retrieved from environment variables.
func NewConfig() *Config {
	smtpPortStr := getEnv("SMTP_PORT", "587")
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		smtpPort = 587
	}
	return &Config{
		EmailSender:    getEnv("EMAIL_SENDER", ""),
		EmailPassword:  getEnv("EMAIL_PASSWORD", ""),
		EmailRecipient: getEnv("EMAIL_RECIPIENT", ""),
		SMTPHost:       getEnv("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:       smtpPort,
	}
}

// Retrieves the value of the specified environment variable, or returns the provided default value if it's not set.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
