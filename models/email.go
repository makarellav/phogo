package models

import (
	"fmt"
	"github.com/go-mail/mail"
)

const (
	DefaultSender = "support@phogo.com"
)

type EmailService struct {
	DefaultSender string
	dialer        *mail.Dialer
}

type Email struct {
	From      string
	To        string
	Subject   string
	Plaintext string
	HTML      string
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewEmailService(config SMTPConfig) *EmailService {
	return &EmailService{
		dialer: mail.NewDialer(config.Host, config.Port, config.Username, config.Password),
	}
}

func (es *EmailService) Send(email Email) error {
	msg := mail.NewMessage()
	msg.SetHeader("To", email.To)
	es.setFrom(msg, email)
	msg.SetHeader("Subject", email.Subject)

	switch {
	case email.HTML != "" && email.Plaintext != "":
		msg.SetBody("text/plain", email.Plaintext)
		msg.AddAlternative("text/html", email.HTML)
	case email.HTML != "":
		msg.SetBody("text/html", email.HTML)
	case email.Plaintext != "":
		msg.SetBody("text/plain", email.Plaintext)
	}

	err := es.dialer.DialAndSend(msg)

	if err != nil {
		return fmt.Errorf("send: %w", err)
	}

	return nil
}

func (es *EmailService) ForgotPassword(to, resetURL string) error {
	err := es.Send(Email{
		To:        to,
		Subject:   "Reset your password",
		Plaintext: fmt.Sprintf("To reset your password, please visit the following link: %s", resetURL),
		HTML:      fmt.Sprintf("<p>To reset your password, please visit the following link: <a href='%s'>%s</a></p>", resetURL, resetURL),
	})

	if err != nil {
		return fmt.Errorf("reset password email: %w", err)
	}

	return nil
}

func (es *EmailService) setFrom(msg *mail.Message, email Email) {
	var from string

	switch {
	case email.From != "":
		from = email.From
	case es.DefaultSender != "":
		from = es.DefaultSender
	default:
		from = DefaultSender
	}

	msg.SetHeader("From", from)
}
