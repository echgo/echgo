package smtp

import "net/mail"

// Access is to config the smtp access data
type Access struct {
	Host     string
	Port     int
	Username mail.Address
	Password string
}

// Data is to structure the email data
type Data struct {
	Email   mail.Address
	Subject string
	Type    string
	Message string
}
