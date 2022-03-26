package email

import (
	"github.com/echgo/echgo/configuration"
	"log"
)

// Execute is to execute the send function
// & lead all configuration data
func Execute(headline, message string) {

	s := Smtp{
		Host:     configuration.Data.Channels.SMTP.Host,
		Port:     configuration.Data.Channels.SMTP.Port,
		Username: configuration.Data.Channels.SMTP.Username,
		Password: configuration.Data.Channels.SMTP.Password,
	}

	d := Data{
		Email:   configuration.Data.Channels.SMTP.EmailRecipient,
		Subject: headline,
		File:    nil,
		HTML:    message,
	}

	err := Send(d, s)
	if err != nil {
		log.Fatalln(err)
	}

}
