package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"strconv"
)

// Smtp is to config the email data
type Smtp struct {
	Host     string
	Port     string
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
func Send(data Data, s Smtp) error {

	m := gomail.NewMessage()

	m.SetHeader("From", s.Username)
	m.SetHeader("To", data.Email)
	m.SetHeader("Subject", data.Subject)

	if data.File != nil {
		m.Attach(*data.File)
	}

	m.SetBody("text/html", data.HTML)

	port, err := strconv.Atoi(s.Port)
	if err != nil {
		return err
	}

	d := gomail.NewDialer(s.Host, port, s.Username, s.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err = d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil

}
