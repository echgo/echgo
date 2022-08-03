package smtp

import (
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/environment"
	"net/mail"
)

// Execute is to execute the send function
// & load all configuration data
func Execute(headline, message string) {

	d := Data{
		Email: mail.Address{
			Address: environment.String("SMTP_EMAIL_RECIPIENT"),
		},
		Subject: headline,
		Type:    "text/plain",
		Message: message,
	}

	a := Access{
		Host: environment.String("SMTP_HOST"),
		Port: environment.Integer("SMTP_PORT"),
		Username: mail.Address{
			Address: environment.String("SMTP_USERNAME"),
		},
		Password: environment.String("SMTP_PASSWORD"),
	}

	err := SendEmail(d, a)
	if err != nil {
		attributes := make(map[string]any)
		attributes["error"] = err
		console.Log("error", "An error occurred while sending the email via smtp.", attributes)
	}

}
