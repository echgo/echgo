package pushover

import (
	"fmt"
	"github.com/echgo/echgo/configuration"
	"log"
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
		log.Fatalln(err)
	}

}
