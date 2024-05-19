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

	if err := smtp.SendMail(m.SmtpHost+":"+m.SmtpPort, m.Auth, m.From, []string{to}, []byte(msg)); err != nil {
		return err
	}

	return nil
}
