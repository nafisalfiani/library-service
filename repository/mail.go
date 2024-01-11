package repository

import (
	"library/entity"

	"gopkg.in/gomail.v2"
)

type mail struct {
	dialer *gomail.Dialer
}

type MailInterface interface {
	Send(header entity.Mail) error
}

// initMail create mail repository
func initMail(dialer *gomail.Dialer) MailInterface {
	return &mail{
		dialer: dialer,
	}
}

func (m *mail) Send(header entity.Mail) error {
	newMail := gomail.NewMessage()
	newMail.SetHeader("From", header.From)
	newMail.SetHeader("To", header.To)
	newMail.SetHeader("Subject", header.Subject)
	newMail.SetBody("text/html", header.Body)

	if err := m.dialer.DialAndSend(newMail); err != nil {
		return err
	}

	return nil
}
