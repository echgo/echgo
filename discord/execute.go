// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package discord

import (
	"fmt"
	"github.com/echgo/echgo/v2/console"
	"github.com/echgo/echgo/v2/environment"
)

// Execute is to execute the create message function
// & lead all configuration data.
func Execute(headline, message string) {

	lookup := environment.Lookup("DISCORD_WEBHOOK_URL")
	if lookup {

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
