package osticket

import (
	"fmt"
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
)

// Execute is to execute the create ticket function
// & lead all configuration data
func Execute(headline, message string) {

	r := Request{
		BaseUrl:  configuration.Data.Channels.OsTicket.BaseUrl,
		ApiToken: configuration.Data.Channels.OsTicket.ApiToken,
	}

	b := CreateTicketBody{
		Alert:       true,
		Autorespond: true,
		Source:      "API",
		Name:        configuration.Data.Channels.OsTicket.Username,
		Email:       configuration.Data.Channels.OsTicket.EmailRecipient,
		Phone:       "",
		Subject:     headline,
		Ip:          "",
		Message:     fmt.Sprintf("data:text/plain,%s", message),
		Attachments: nil,
	}

	_, err := CreateTicket(b, r)
	if err != nil {
		attributes := make(map[string]any)
		attributes["error"] = err
		console.Log("error", "An error occurred while creating the ticket via osticket.", attributes)
	}

}
