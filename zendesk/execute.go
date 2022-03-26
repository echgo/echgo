package zendesk

import (
	"github.com/echgo/echgo/configuration"
	"log"
)

// Execute is to execute the create ticket function
// & lead all configuration data
func Execute(headline, message string) {

	r := Request{
		BaseUrl:  configuration.Data.Channels.Zendesk.BaseUrl,
		Username: configuration.Data.Channels.Zendesk.Username,
		ApiToken: configuration.Data.Channels.Zendesk.ApiToken,
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
		log.Fatalln(err)
	}

}
