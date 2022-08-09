package twillo

import (
	"fmt"
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/environment"
)

// Execute is to execute to send message function
// & lead all configuration data
func Execute(headline, message string) {

	lookup := environment.Lookup("TWILLO_ACCOUNT_SID", "TWILLO_AUTH_TOKEN", "TWILLO_PHONE_NUMBER", "TWILLO_MY_PHONE_NUMBER")
	if lookup {

		r := Request{
			AccountSid: environment.String("TWILLO_ACCOUNT_SID"),
			AuthToken:  environment.String("TWILLO_AUTH_TOKEN"),
		}

		b := CreateMessageBody{
			Message:       fmt.Sprintf("%s - %s", headline, message),
			PhoneNumber:   environment.String("TWILLO_PHONE_NUMBER"),
			MyPhoneNumber: environment.String("TWILLO_MY_PHONE_NUMBER"),
		}

		_, err := CreateMessage(b, r)
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
