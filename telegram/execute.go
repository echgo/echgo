package telegram

import (
	"fmt"
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
)

// Execute is to execute the create message function
// & lead all configuration data
func Execute(headline, message string) {

	r := Request{
		ApiToken: configuration.Data.Channels.Telegram.ApiToken,
		ChatId:   configuration.Data.Channels.Telegram.ChatId,
	}

	_, err := CreateMessage(fmt.Sprintf("%s\n%s", headline, message), "Markdown", r)
	if err != nil {
		attributes := make(map[string]any)
		attributes["error"] = err
		console.Log("error", "An error occurred while creating the message via telegram.", attributes)
	}

}
