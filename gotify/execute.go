package gotify

import (
	"github.com/echgo/echgo/configuration"
	"log"
)

// Execute is to execute the create message function
// & lead all configuration data
func Execute(headline, message string) {

	r := Request{
		Domain:     configuration.Data.Channels.Gotify.Domain,
		XGotifyKey: configuration.Data.Channels.Gotify.Key,
	}

	b := CreateMessageBody{
		Priority: 0,
		Title:    headline,
		Message:  message,
		Extras:   nil,
	}

	_, err := CreateMessage(b, r)
	if err != nil {
		log.Fatalln(err)
	}

}
