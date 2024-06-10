package mailer

import (
	"net/smtp"
	"server/internal/config"
)

type Mailer struct {
	From     string
	Password string
	SmtpHost string
	SmtpPort string
	Auth     smtp.Auth
}

func NewMailer(cfg config.Mailer) *Mailer {
	auth := smtp.PlainAuth("", cfg.From, cfg.Password, cfg.SmtpHost)

	return &Mailer{
		From:     cfg.From,
		Password: cfg.Password,
		SmtpPort: cfg.SmtpPort,
		SmtpHost: cfg.SmtpHost,
		Auth:     auth,
	}
}

func (m Mailer) SendMessage(to, msg string) error {

	message := "From: " + m.From + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: NO REPLY (Finkomek fintech app)\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\";\r\n" +
		"\r\n" + msg

	if err := smtp.SendMail(m.SmtpHost+":"+m.SmtpPort, m.Auth, m.From, []string{to}, []byte(message)); err != nil {
		return err
	}

	return nil
}
