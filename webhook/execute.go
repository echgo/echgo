// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package webhook

import (
	"github.com/echgo/echgo/v2/console"
	"github.com/echgo/echgo/v2/environment"
)

// Execute is to execute the webhook
// function & lead all configuration data.
func Execute(headline, message string) {

	lookup := environment.Lookup("WEBHOOK_URL")
	if lookup {

		r := Request{
			Url: environment.String("WEBHOOK_URL"),
		}

		b := SendBody{
			Headline: headline,
			Message:  message,
		}

		err := Send(b, r)
		if err != nil {

			attributes := make(map[string]any)
			attributes["channel"] = channel
			attributes["error"] = err
			console.Log("error", "An error occurred while send the data.", attributes)

		}

	} else {

		attributes := make(map[string]any)
		attributes["channel"] = channel
		attributes["lookup"] = lookup
		console.Log("error", "An error occurred while lookup the environment variables.", attributes)

	}

}
