package services

import (
	"fmt"
	"net/smtp"
)

// Gmail SMTP credentials (replace with your own).
const (
	gmailSMTPServer = "smtp.gmail.com"
	gmailSMTPPort   = "587"
	gmailUsername   = "your-email@gmail.com"
	gmailPassword   = "your-email-password"
)

func SendWelcomeEmail(email string) error {
	// Set up the email content.
	subject := "Welcome to Our App"
	body := fmt.Sprintf("Hello, %s! Welcome to our app. We're glad to have you.", email)
	message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

	// Set up authentication information.
	auth := smtp.PlainAuth("", gmailUsername, gmailPassword, gmailSMTPServer)

	// Send the email.
	err := smtp.SendMail(gmailSMTPServer+":"+gmailSMTPPort, auth, gmailUsername, []string{email}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
