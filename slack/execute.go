package slack

import (
	"fmt"
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
)

// Execute is to execute the create message function
// & lead all configuration data
func Execute(headline, message string) {

	r := Request{
		Url: configuration.Data.Channels.Slack.WebhookUrl,
	}

	b := CreateMessageBody{
		Text: fmt.Sprintf("%s - %s", headline, message),
	}

	_, err := CreateMessage(b, r)
	if err != nil {
		attributes := make(map[string]any)
		attributes["error"] = err
		console.Log("error", "An error occurred while creating the message via slack.", attributes)
	}

}
