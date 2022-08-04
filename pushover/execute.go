package pushover

import (
	"fmt"
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/environment"
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
			attributes["error"] = err
			console.Log("error", "An error occurred while creating the message via pushover.", attributes)
		}

	}

}
