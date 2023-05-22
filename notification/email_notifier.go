package notification

import (
	"fmt"
	"github.com/sendinblue/httpsms-go"
)

// EmailNotifier represents an email notification service.
type EmailNotifier struct {
	SenderEmail   string
	RecipientEmail string
	EmailService  string
}

// NewEmailNotifier creates a new email notifier with the specified email address and service.
func NewEmailNotifier(senderEmail, recipientEmail, emailService string) *EmailNotifier {
	return &EmailNotifier{
		SenderEmail:   senderEmail,
		RecipientEmail: recipientEmail,
		EmailService:  emailService,
	}
}

// SendNotification sends an email notification with the specified message.
func (n *EmailNotifier) SendNotification(message string) error {
	// Implement your logic to send an email notification using the configured sender, recipient, and email service
	emailClient := httpsms.NewSmtpClient(n.EmailService)
	emailClient.Host = "smtp-relay.sendinblue.com"
	emailClient.Port = 587
	emailClient.Username = "aud1sm0@live.com"
	emailClient.Password = "Player2loggedin"
	emailClient.Tls = true

	email := &httpsms.EmailMessage{
		From:    n.SenderEmail,
		To:      []string{n.RecipientEmail},
		Subject: "URL Monitoring Alert",
		Text:    message,
	}

	err := emailClient.Send(email)
	if err != nil {
		return fmt.Errorf("failed to send email notification: %v", err)
	}

	return nil
}
