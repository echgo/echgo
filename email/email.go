package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

// Smtp is to config the email data
type Smtp struct {
	Host     string
	Port     int
	Username string
	Password string
}

// Data is to structure the data
type Data struct {
	Email   string
	Subject string
	File    *string
	HTML    string
}

// Send is to send an email via email
// You can send the emails with smtp & add attachment files
func Send(d Data, s Smtp) error {

	m := gomail.NewMessage()

	m.SetHeader("From", s.Username)
	m.SetHeader("To", d.Email)
	m.SetHeader("Subject", d.Subject)

	if d.File != nil {
		m.Attach(*d.File)
	}

	m.SetBody("text/html", d.HTML)

	dialer := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := dialer.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil

}
