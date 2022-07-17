package webhook

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
)

// Execute is to execute the webhook function
// & lead all configuration data
func Execute(headline, message string) {

	r := Request{
		Domain: configuration.Data.Channels.Webhook.Domain,
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

}
