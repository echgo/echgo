// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package trello

import (
	"github.com/echgo/echgo/v2/console"
	"github.com/echgo/echgo/v2/environment"
)

// Execute is to execute the create card
// function & lead all configuration data.
func Execute(headline, message string) {

	lookup := environment.Lookup("TRELLO_KEY", "TRELLO_TOKEN", "TRELLO_ID_LIST")
	if lookup {

		r := Request{
			Key:    environment.String("TRELLO_KEY"),
			Token:  environment.String("TRELLO_TOKEN"),
			IdList: environment.String("TRELLO_ID_LIST"),
		}

		_, err := CreateCard(headline, message, r)
		if err != nil {

			attributes := make(map[string]any)
			attributes["channel"] = channel
			attributes["error"] = err
			console.Log("error", "An error occurred while creating the card.", attributes)

		}

	} else {

		attributes := make(map[string]any)
		attributes["channel"] = channel
		attributes["lookup"] = lookup
		console.Log("error", "An error occurred while lookup the environment variables.", attributes)

	}

}
