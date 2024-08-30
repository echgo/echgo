// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package smtp is used to be able to send an email via smtp.
package smtp

import "net/mail"

// channel is to save the channel name for logging.
const channel = "smtp"

// Access is to config the smtp access data.
type Access struct {
	Host     string
	Port     int
	Username mail.Address
	Password string
}

// Data is to structure the email data.
type Data struct {
	Email   mail.Address
	Subject string
	Type    string
	Message string
}
