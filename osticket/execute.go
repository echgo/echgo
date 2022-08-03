package osticket

import (
	"fmt"
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/environment"
)

// Execute is to execute the create ticket function
// & lead all configuration data
func Execute(headline, message string) {

	r := Request{
		BaseUrl:  environment.String("OS_TICKET_BASE_URL"),
		ApiToken: environment.String("OS_TICKET_API_TOKEN"),
	}

	b := CreateTicketBody{
		Alert:       true,
		Autorespond: true,
		Source:      "API",
		Name:        environment.String("OS_TICKET_USERNAME"),
		Email:       environment.String("OS_TICKET_EMAIL_RECIPIENT"),
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
