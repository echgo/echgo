package matrix

import (
	"fmt"
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
)

// Execute is to execute the create message function
// & lead all configuration data
func Execute(headline, message string) {

	r := Request{
		Domain:      configuration.Data.Channels.Matrix.Domain,
		RoomId:      configuration.Data.Channels.Matrix.RoomId,
		AccessToken: configuration.Data.Channels.Matrix.AccessToken,
	}

	b := CreateMessageBody{
		Msgtype: "m.text",
		Body:    fmt.Sprintf("%s\n%s", headline, message),
	}

	_, err := CreateMessage(b, r)
	if err != nil {
		attributes := make(map[string]any)
		attributes["error"] = err
		console.Log("error", "An error occurred while creating the message via matrix.", attributes)
	}

}
