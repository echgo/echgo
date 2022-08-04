package matrix

import (
	"fmt"
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/environment"
)

// Execute is to execute the create message function
// & lead all configuration data
func Execute(headline, message string) {

	lookup := environment.Lookup("MATRIX_DOMAIN", "MATRIX_ROOM_ID", "MATRIX_ACCESS_TOKEN")
	if lookup {

		r := Request{
			Domain:      environment.String("MATRIX_DOMAIN"),
			RoomId:      environment.String("MATRIX_ROOM_ID"),
			AccessToken: environment.String("MATRIX_ACCESS_TOKEN"),
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

}
