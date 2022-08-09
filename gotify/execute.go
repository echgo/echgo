package gotify

import (
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/environment"
)

// Execute is to execute the create message function
// & lead all configuration data
func Execute(headline, message string) {

	lookup := environment.Lookup("GOTIFY_BASE_URL", "GOTIFY_KEY")
	if lookup {

		r := Request{
			BaseUrl:    environment.String("GOTIFY_BASE_URL"),
			XGotifyKey: environment.String("GOTIFY_KEY"),
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
			attributes["channel"] = "gotify"
			attributes["error"] = err
			console.Log("error", "An error occurred while creating the message.", attributes)

		}

	} else {

		attributes := make(map[string]any)
		attributes["channel"] = "gotify"
		attributes["lookup"] = lookup
		console.Log("error", "An error occurred while lookup the environment variables.", attributes)

	}

}
