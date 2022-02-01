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
func Send(data Data, s Smtp) error {

	// Config go mail
	m := gomail.NewMessage()

	// Set header
	m.SetHeader("From", s.Username)
	m.SetHeader("To", data.Email)
	m.SetHeader("Subject", data.Subject)

	// Check if file exists & add file
	if data.File != nil {
		m.Attach(*data.File)
	}

	// Set body
	m.SetBody("text/html", data.HTML)

	// Config dialer
	d := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send email
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	// Return none error
	return nil

}
