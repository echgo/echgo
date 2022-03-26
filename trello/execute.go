package trello

import (
	"github.com/echgo/echgo/configuration"
	"log"
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
		log.Fatalln(err)
	}

}
