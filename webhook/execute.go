package webhook

import (
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/environment"
)

// Execute is to execute the webhook function
// & lead all configuration data
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
			attributes["error"] = err
			console.Log("error", "An error occurred while sending the webhook.", attributes)
		}

	} else {

		attributes := make(map[string]any)
		attributes["channel"] = "webhook"
		attributes["lookup"] = lookup
		console.Log("error", "An error occurred while lookup the environment variables.", attributes)

	}

}
