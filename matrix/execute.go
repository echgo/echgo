package matrix

import (
	"fmt"
	"github.com/echgo/echgo/v2/console"
	"github.com/echgo/echgo/v2/environment"
)

// Execute is to execute the create message function
// & lead all configuration data
func Execute(headline, message string) {

	lookup := environment.Lookup("MATRIX_BASE_URL", "MATRIX_ROOM_ID", "MATRIX_ACCESS_TOKEN")
	if lookup {

		r := Request{
			BaseUrl:     environment.String("MATRIX_BASE_URL"),
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
