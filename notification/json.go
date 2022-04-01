package notification

import (
	"encoding/json"
	"github.com/echgo/echgo/channels"
	"io"
	"log"
	"os"
)

// JsonBody is to decode the json file
type JsonBody struct {
	Channels []string `json:"channels"`
	Headline string   `json:"headline"`
	Message  string   `json:"message"`
}

// Json is to handle the json files from the notification
// Check them & send them to the channel handler
func Json(file *os.File) {

	var decode JsonBody

	read, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(read, &decode)
	if err != nil {
		log.Fatalln(err)
	}

	var channel []string
	for _, value := range decode.Channels {
		channel = append(channel, value)
	}

	channels.Handler(decode.Headline, decode.Message, channel)

}
