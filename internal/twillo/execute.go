// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package twillo

import (
	"fmt"
	"github.com/echgo/echgo/v2/internal/console"
	environment2 "github.com/echgo/echgo/v2/internal/environment"
)

// Execute is to execute to send message
// function & lead all configuration data.
func Execute(headline, message string) {

	lookup := environment2.Lookup("TWILLO_ACCOUNT_SID", "TWILLO_AUTH_TOKEN", "TWILLO_PHONE_NUMBER", "TWILLO_MY_PHONE_NUMBER")
	if lookup {

		r := Request{
			AccountSid: environment2.String("TWILLO_ACCOUNT_SID"),
			AuthToken:  environment2.String("TWILLO_AUTH_TOKEN"),
		}

		b := CreateMessageBody{
			Message:       fmt.Sprintf("%s - %s", headline, message),
			PhoneNumber:   environment2.String("TWILLO_PHONE_NUMBER"),
			MyPhoneNumber: environment2.String("TWILLO_MY_PHONE_NUMBER"),
		}

		_, err := CreateMessage(b, r)
		if err != nil {

			attributes := make(map[string]any)
			attributes["channel"] = channel
			attributes["error"] = err
			console.Log("error", "An error occurred while creating the message.", attributes)

		}

	} else {

		attributes := make(map[string]any)
		attributes["channel"] = channel
		attributes["lookup"] = lookup
		console.Log("error", "An error occurred while lookup the environment variables.", attributes)

	}

}
