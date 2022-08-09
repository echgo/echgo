package zendesk

import (
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/environment"
)

// Execute is to execute the create ticket function
// & lead all configuration data
func Execute(headline, message string) {

	lookup := environment.Lookup("ZENDESK_BASE_URL", "ZENDESK_USERNAME", "ZENDESK_API_TOKEN")
	if lookup {

		r := Request{
			BaseUrl:  environment.String("ZENDESK_BASE_URL"),
			Username: environment.String("ZENDESK_USERNAME"),
			ApiToken: environment.String("ZENDESK_API_TOKEN"),
		}

		b := CreateTicketBody{
			Ticket: CreateTicketBodyTicket{
				Comment: CreateTicketBodyComment{
					Body: message,
				},
				Priority: "urgent",
				Subject:  headline,
			},
		}

		_, err := CreateTicket(b, r)
		if err != nil {

			attributes := make(map[string]any)
			attributes["channel"] = "zendesk"
			attributes["error"] = err
			console.Log("error", "An error occurred while creating the ticket.", attributes)

		}

	} else {

		attributes := make(map[string]any)
		attributes["channel"] = "zendesk"
		attributes["lookup"] = lookup
		console.Log("error", "An error occurred while lookup the environment variables.", attributes)

	}

}
