// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package osticket

import (
	"fmt"
	"github.com/echgo/echgo/v2/console"
	"github.com/echgo/echgo/v2/environment"
)

// Execute is to execute the create ticket
// function & lead all configuration data.
func Execute(headline, message string) {

	lookup := environment.Lookup("OS_TICKET_BASE_URL", "OS_TICKET_API_TOKEN", "OS_TICKET_USERNAME", "OS_TICKET_EMAIL_RECIPIENT")
	if lookup {

		r := Request{
			BaseUrl:  environment.String("OS_TICKET_BASE_URL"),
			ApiToken: environment.String("OS_TICKET_API_TOKEN"),
		}

		b := CreateTicketBody{
			Alert:       true,
			Autorespond: true,
			Source:      "API",
			Name:        environment.String("OS_TICKET_USERNAME"),
			Email:       environment.String("OS_TICKET_EMAIL_RECIPIENT"),
			Phone:       "",
			Subject:     headline,
			Ip:          "",
			Message:     fmt.Sprintf("data:text/plain,%s", message),
			Attachments: nil,
		}

		_, err := CreateTicket(b, r)
		if err != nil {

			attributes := make(map[string]any)
			attributes["channel"] = channel
			attributes["error"] = err
			console.Log("error", "An error occurred while creating the ticket.", attributes)

		}

	} else {

		attributes := make(map[string]any)
		attributes["channel"] = channel
		attributes["lookup"] = lookup
		console.Log("error", "An error occurred while lookup the environment variables.", attributes)

	}

}
