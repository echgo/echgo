package trello

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
)

// Execute is to execute the create card function
// & lead all configuration data
func Execute(headline, message string) {

	r := Request{
		Key:    configuration.Data.Channels.Trello.Key,
		Token:  configuration.Data.Channels.Trello.Token,
		IdList: configuration.Data.Channels.Trello.IdList,
	}

	_, err := CreateCard(headline, message, r)
	if err != nil {
		attributes := make(map[string]any)
		attributes["error"] = err
		console.Log("error", "An error occurred while creating the card via trello.", attributes)
	}

}
