package smtputil

import (
	"context"
	"learnyscape-backend-mono/pkg/config"

	"gopkg.in/gomail.v2"
)

type Mailer interface {
	SendMail(ctx context.Context, to, subject, body string) error
}

type mailerImpl struct {
	dialer *gomail.Dialer
}

func NewMailer() Mailer {
	return &mailerImpl{
		dialer: gomail.NewDialer(
			config.SmtpConfig.Host,
			config.SmtpConfig.Port,
			"",
			"",
		),
	}
}

func (m *mailerImpl) SendMail(ctx context.Context, to, subject, body string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		msg := gomail.NewMessage()

		msg.SetHeader("From", config.SmtpConfig.Email)
		msg.SetHeader("To", to)
		msg.SetHeader("Subject", subject)
		msg.SetBody("text/html", body)

		return m.dialer.DialAndSend(msg)
	}
}
