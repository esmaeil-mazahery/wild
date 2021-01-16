package email

import (
	"context"
	"net/smtp"
)

// IProvider is an interface to sms provider
type IProvider interface {
	//Send save Service data to the store
	Send(ctx context.Context, message string, receptor []string) error
}

// Provider stores users in memory
type Provider struct {
	from     string
	password string
	smtpHost string
	smtpPort string
}

// NewProvider returns a new in-memory user store
func NewProvider(
	from string,
	password string,
	smtpHost string,
	smtpPort string,
) *Provider {
	return &Provider{
		from:     from,
		password: password,
		smtpHost: smtpHost,
		smtpPort: smtpPort,
	}
}

//Send send email to Provider
func (p *Provider) Send(ctx context.Context, message string, receptor []string) error {

	// Message.
	messageByte := []byte(message)

	// Authentication.
	auth := smtp.PlainAuth("", p.from, p.password, p.smtpHost)

	// Sending email.
	err := smtp.SendMail(p.smtpHost+":"+p.smtpPort, auth, p.from, receptor, messageByte)
	if err != nil {

		return err
	}
	return nil
}
