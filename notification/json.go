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
	Channels struct {
		Gotify   bool `json:"gotify,omitempty"`
		Matrix   bool `json:"matrix,omitempty"`
		Telegram bool `json:"telegram,omitempty"`
		Discord  bool `json:"discord,omitempty"`
		Slack    bool `json:"slack,omitempty"`
		Trello   bool `json:"trello,omitempty"`
		Zendesk  bool `json:"zendesk,omitempty"`
		Osticket bool `json:"osticket,omitempty"`
		Smtp     bool `json:"smtp,omitempty"`
		Webhook  bool `json:"webhook,omitempty"`
	} `json:"channels"`
	Headline string `json:"headline"`
	Message  string `json:"message"`
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

	channel := channels.Type{
		Gotify:   decode.Channels.Gotify,
		Matrix:   decode.Channels.Matrix,
		Telegram: decode.Channels.Telegram,
		Discord:  decode.Channels.Discord,
		Slack:    decode.Channels.Slack,
		Trello:   decode.Channels.Trello,
		Zendesk:  decode.Channels.Zendesk,
		Osticket: decode.Channels.Osticket,
		SMTP:     decode.Channels.Smtp,
		Webhook:  decode.Channels.Webhook,
	}

	channels.Handler(decode.Headline, decode.Message, channel)

}
