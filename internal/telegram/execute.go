// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package telegram

import (
	"fmt"
	"github.com/echgo/echgo/v2/internal/console"
	environment2 "github.com/echgo/echgo/v2/internal/environment"
)

// Execute is to execute the create message
// function & lead all configuration data.
func Execute(headline, message string) {

	lookup := environment2.Lookup("TELEGRAM_API_TOKEN", "TELEGRAM_CHAT_ID")
	if lookup {

		r := Request{
			ApiToken: environment2.String("TELEGRAM_API_TOKEN"),
			ChatId:   environment2.String("TELEGRAM_CHAT_ID"),
		}

		_, err := CreateMessage(fmt.Sprintf("%s\n%s", headline, message), "Markdown", r)
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
