package notification

import (
	"encoding/json"
	"github.com/echgo/echgo/channels"
	"io"
	"log"
	"os"
)

// JSONBody is to decode the json file
type JSONBody struct {
	Channels struct {
		Gotify   bool `json:"gotify,omitempty"`
		Matrix   bool `json:"matrix,omitempty"`
		Telegram bool `json:"telegram,omitempty"`
		Discord  bool `json:"discord,omitempty"`
		Trello   bool `json:"trello,omitempty"`
		Smtp     bool `json:"smtp,omitempty"`
		Webhook  bool `json:"webhook,omitempty"`
	} `json:"channels"`
	Headline string `json:"headline"`
	Message  string `json:"message"`
}

// JSON is to handle the json files from the notification
// Check them & send them to the channel handler
// Them remove the files
func JSON(path string) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var decode JSONBody

	read, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(read, &decode)
	if err != nil {
		log.Fatalln(err)
	}

	channel := channels.Type{
		Gotify:   decode.Channels.Gotify,
		Matrix:   decode.Channels.Matrix,
		Telegram: decode.Channels.Telegram,
		Discord:  decode.Channels.Discord,
		Trello:   decode.Channels.Trello,
		SMTP:     decode.Channels.Smtp,
		Webhook:  decode.Channels.Webhook,
	}

	channels.Handler(decode.Headline, decode.Message, channel)

	err = os.Remove(path)
	if err != nil {
		log.Fatalln(err)
	}

}
