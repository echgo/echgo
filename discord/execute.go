package discord

import (
	"fmt"
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
)

// Execute is to execute the create message function
// & lead all configuration data
func Execute(headline, message string) {

	r := Request{
		Url: configuration.Data.Channels.Discord.WebhookUrl,
	}

	if len(configuration.Data.Channels.Discord.BotName) == 0 {
		configuration.Data.Channels.Discord.BotName = "echGo"
	}

	b := CreateMessageBody{
		Username:  configuration.Data.Channels.Discord.BotName,
		AvatarUrl: "https://raw.githubusercontent.com/echgo/logo/main/logo-small.png",
		Content:   fmt.Sprintf("%s - %s", headline, message),
	}

	err := CreateMessage(b, r)
	if err != nil {
		attributes := make(map[string]any)
		attributes["error"] = err
		console.Log("error", "An error occurred while creating the message via discord.", attributes)
	}

}
