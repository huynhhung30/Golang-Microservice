package utils

import (
	"net/smtp"
	"os"
	"user_service/utils/functions"
)

func SendMail(email string, subject string, body string) {
	emailHost := os.Getenv("EMAIL_HOST")
	emailHostFrom := os.Getenv("EMAIL_HOST_FROM")
	emailHostUser := os.Getenv("EMAIL_HOST_USER")
	emailHostPassword := os.Getenv("EMAIL_HOST_PASSWORD")

	// Sender data.
	from := emailHostFrom
	user := emailHostUser
	password := emailHostPassword

	msg := []byte("From: " + from + "\r\n" +
		"To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body + "\r\n")

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := emailHost
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", user, password, smtpHost)
	// messageBody := []byte(message)
	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)

	// Now send E-Mail
	if err != nil {
		functions.ShowLog("SendMailError", err)
	}
}
