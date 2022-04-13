package twillo

import (
	"fmt"
	"github.com/echgo/echgo/configuration"
	"log"
)

// Execute is to execute the send message function
// & lead all configuration data
func Execute(headline, message string) {

	r := Request{
		AccountSid: configuration.Data.Channels.Twillo.AccountSid,
		AuthToken:  configuration.Data.Channels.Twillo.AuthToken,
	}

	b := CreateMessageBody{
		Message:       fmt.Sprintf("%s - %s", headline, message),
		PhoneNumber:   configuration.Data.Channels.Twillo.PhoneNumber,
		MyPhoneNumber: configuration.Data.Channels.Twillo.MyPhoneNumber,
	}

	_, err := CreateMessage(b, r)
	if err != nil {
		log.Fatalln(err)
	}

}
