// Copyright 2024 Jonas Kwiedor. All rights reserved.

package gotify

import (
	"github.com/echgo/echgo/v2/internal/console"
	environment2 "github.com/echgo/echgo/v2/internal/environment"
)

// Execute is to execute the create message function & lead all configuration data.
func Execute(headline, message string) {

	lookup := environment2.Lookup("GOTIFY_BASE_URL", "GOTIFY_KEY")
	if lookup {

		r := Request{
			BaseUrl:    environment2.String("GOTIFY_BASE_URL"),
			XGotifyKey: environment2.String("GOTIFY_KEY"),
		}

		b := CreateMessageBody{
			Priority: 0,
			Title:    headline,
			Message:  message,
			Extras:   nil,
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
