package smtp

import (
	"github.com/echgo/echgo/configuration"
	"log"
	"net/mail"
)

// Execute is to execute the send function
// & load all configuration data
func Execute(headline, message string) {

	d := Data{
		Email: mail.Address{
			Address: configuration.Data.Channels.SMTP.EmailRecipient,
		},
		Subject: headline,
		Type:    "text/plain",
		Message: message,
	}

	a := Access{
		Host: configuration.Data.Channels.SMTP.Host,
		Port: configuration.Data.Channels.SMTP.Port,
		Username: mail.Address{
			Address: configuration.Data.Channels.SMTP.Username,
		},
		Password: configuration.Data.Channels.SMTP.Password,
	}

	err := SendEmail(d, a)
	if err != nil {
		log.Fatalln(err)
	}

}
