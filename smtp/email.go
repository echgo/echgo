// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package smtp

import (
	"crypto/tls"
	"net/smtp"
	"strconv"
)

// SendEmail is to send an email. You can send
// the emails with smtp & add attachment files
func SendEmail(d Data, a Access) error {

	headers := make(map[string]string)
	headers["Content-Type"] = d.Type
	headers["From"] = a.Username.String()
	headers["To"] = d.Email.String()
	headers["Subject"] = d.Subject

	var message string
	for index, value := range headers {
		message += index + ": " + value + "\r\n"
	}
	message += "\r\n" + d.Message

	authentication := smtp.PlainAuth("", a.Username.Address, a.Password, a.Host)

	configuration := &tls.Config{
		ServerName: a.Host,
	}

	connection, err := tls.Dial("tcp", a.Host+":"+strconv.Itoa(a.Port), configuration)
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(connection, a.Host)
	if err != nil {
		return err
	}

	err = client.Auth(authentication)
	if err != nil {
		return err
	}

	err = client.Mail(a.Username.Address)
	if err != nil {
		return err
	}

	err = client.Rcpt(d.Email.Address)
	if err != nil {
		return err
	}

	writer, err := client.Data()
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(message))
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	err = client.Quit()
	if err != nil {
		return err
	}

	return nil

}
