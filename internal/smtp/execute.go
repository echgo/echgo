// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package smtp

import (
	"github.com/echgo/echgo/v2/internal/console"
	environment2 "github.com/echgo/echgo/v2/internal/environment"
	"net/mail"
)

// Execute is to execute the send
// function & load all configuration data
func Execute(headline, message string) {

	lookup := environment2.Lookup("SMTP_HOST", "SMTP_PORT", "SMTP_USERNAME", "SMTP_PASSWORD", "SMTP_EMAIL_RECIPIENT")
	if lookup {

		d := Data{
			Email: mail.Address{
				Address: environment2.String("SMTP_EMAIL_RECIPIENT"),
			},
			Subject: headline,
			Type:    "text/plain",
			Message: message,
		}

		a := Access{
			Host: environment2.String("SMTP_HOST"),
			Port: environment2.Integer("SMTP_PORT"),
			Username: mail.Address{
				Address: environment2.String("SMTP_USERNAME"),
			},
			Password: environment2.String("SMTP_PASSWORD"),
		}

		err := SendEmail(d, a)
		if err != nil {

			attributes := make(map[string]any)
			attributes["channel"] = channel
			attributes["error"] = err
			console.Log("error", "An error occurred while sending the email.", attributes)

		}

	} else {

		attributes := make(map[string]any)
		attributes["channel"] = channel
		attributes["lookup"] = lookup
		console.Log("error", "An error occurred while lookup the environment variables.", attributes)

	}

}
