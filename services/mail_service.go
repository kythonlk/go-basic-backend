package services

import (
	"fmt"
	"net/smtp"
)

const (
	gmailSMTPServer = "smtp.gmail.com"
	gmailSMTPPort   = "587"
	gmailUsername   = "your-email@gmail.com"
	gmailPassword   = "your-email-password"
)

func SendWelcomeEmail(email string) error {
	subject := "Welcome to Our App"
	body := fmt.Sprintf("Hello, %s! Welcome to our app. We're glad to have you.", email)
	message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

	auth := smtp.PlainAuth("", gmailUsername, gmailPassword, gmailSMTPServer)

	err := smtp.SendMail(gmailSMTPServer+":"+gmailSMTPPort, auth, gmailUsername, []string{email}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
