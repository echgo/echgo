package gotify

import (
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/environment"
)

// Execute is to execute the create message function
// & lead all configuration data
func Execute(headline, message string) {

	r := Request{
		Domain:     environment.String("GOTIFY_DOMAIN"),
		XGotifyKey: environment.String("GOTIFY_KEY"),
	}

	b := CreateMessageBody{
		Priority: 0,
		Title:    headline,
		Message:  message,
		Extras:   nil,
	}

	_, err := CreateMessage(b, r)
	if err != nil {
		attributes := make(map[string]any)
		attributes["error"] = err
		console.Log("error", "An error occurred while creating the message via gotify.", attributes)
	}

}
