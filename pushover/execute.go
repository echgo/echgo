package pushover

import (
	"fmt"
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
)

// Execute is to execute the create message function
// & lead all configuration data
func Execute(headline, message string) {

	b := CreateMessageBody{
		Token:   configuration.Data.Channels.Pushover.Token,
		User:    configuration.Data.Channels.Pushover.User,
		Message: fmt.Sprintf("%s\n%s", headline, message),
	}

	_, err := CreateMessage(b)
	if err != nil {
		attributes := make(map[string]any)
		attributes["error"] = err
		console.Log("error", "An error occurred while creating the message via pushover.", attributes)
	}

}
