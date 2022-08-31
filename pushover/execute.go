package pushover

import (
	"fmt"
	"github.com/echgo/echgo/v2/console"
	"github.com/echgo/echgo/v2/environment"
)

// Execute is to execute the create message function
// & lead all configuration data
func Execute(headline, message string) {

	lookup := environment.Lookup("PUSHOVER_TOKEN", "PUSHOVER_USER")
	if lookup {

		b := CreateMessageBody{
			Token:   environment.String("PUSHOVER_TOKEN"),
			User:    environment.String("PUSHOVER_USER"),
			Message: fmt.Sprintf("%s\n%s", headline, message),
		}

		_, err := CreateMessage(b)
		if err != nil {

			attributes := make(map[string]any)
			attributes["channel"] = channel
			attributes["error"] = err
			console.Log("error", "An error occurred while creating the message.", attributes)

		}

	} else {

		attributes := make(map[string]any)
		attributes["channel"] = channel
		attributes["lookup"] = lookup
		console.Log("error", "An error occurred while lookup the environment variables.", attributes)

	}

}
