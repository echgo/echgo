package discord

import (
	"fmt"
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/environment"
)

// Execute is to execute the create message function
// & lead all configuration data
func Execute(headline, message string) {

	r := Request{
		Url: environment.String("DISCORD_WEBHOOK_URL"),
	}

	b := CreateMessageBody{
		Username:  "echGo",
		AvatarUrl: "https://raw.githubusercontent.com/echgo/logo/main/logo-small.png",
		Content:   fmt.Sprintf("%s - %s", headline, message),
	}

	lookup := environment.Lookup("DISCORD_BOT_NAME")
	if lookup {
		b.Username = environment.String("DISCORD_BOT_NAME")
	}

	err := CreateMessage(b, r)
	if err != nil {
		attributes := make(map[string]any)
		attributes["error"] = err
		console.Log("error", "An error occurred while creating the message via discord.", attributes)
	}

}
