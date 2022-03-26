package smtp

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

// SendEmail is to send an email
// You can send the emails with smtp & add attachment files
func SendEmail(d Data, a Access) error {

	m := gomail.NewMessage()

	m.SetHeader("From", a.Username)
	m.SetHeader("To", d.Email)
	m.SetHeader("Subject", d.Subject)

	if d.File != nil {
		m.Attach(*d.File)
	}

	m.SetBody("text/html", d.HTML)

	dialer := gomail.NewDialer(a.Host, a.Port, a.Username, a.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := dialer.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil

}
