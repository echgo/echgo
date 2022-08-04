package trello

import (
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/environment"
)

// Execute is to execute the create card function
// & lead all configuration data
func Execute(headline, message string) {

	lookup := environment.Lookup("TRELLO_KEY", "TRELLO_TOKEN", "TRELLO_ID_LIST")
	if lookup {

		r := Request{
			Key:    environment.String("TRELLO_KEY"),
			Token:  environment.String("TRELLO_TOKEN"),
			IdList: environment.String("TRELLO_ID_LIST"),
		}

		_, err := CreateCard(headline, message, r)
		if err != nil {
			attributes := make(map[string]any)
			attributes["error"] = err
			console.Log("error", "An error occurred while creating the card via trello.", attributes)
		}

	}

}
