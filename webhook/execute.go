package webhook

import (
	"github.com/echgo/echgo/configuration"
	"log"
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
		log.Fatalln(err)
	}

}
