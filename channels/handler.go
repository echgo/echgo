// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package channels is used to transfer the different headlines
// and messages to the different execute functions of the packages.
package channels

import (
	"github.com/echgo/echgo/v2/discord"
	"github.com/echgo/echgo/v2/gotify"
	"github.com/echgo/echgo/v2/matrix"
	"github.com/echgo/echgo/v2/osticket"
	"github.com/echgo/echgo/v2/pushover"
	"github.com/echgo/echgo/v2/slack"
	"github.com/echgo/echgo/v2/smtp"
	"github.com/echgo/echgo/v2/telegram"
	"github.com/echgo/echgo/v2/trello"
	"github.com/echgo/echgo/v2/twillo"
	"github.com/echgo/echgo/v2/webhook"
	"github.com/echgo/echgo/v2/zendesk"
	"strings"
)

// Handler is to handle the channels and send the notifications to them.
func Handler(headline, message string, types []string) {

	for _, value := range types {
		switch {
		case strings.EqualFold("gotify", value):
			gotify.Execute(headline, message)
		case strings.EqualFold("pushover", value):
			pushover.Execute(headline, message)
		case strings.EqualFold("matrix", value):
			matrix.Execute(headline, message)
		case strings.EqualFold("telegram", value):
			telegram.Execute(headline, message)
		case strings.EqualFold("discord", value):
			discord.Execute(headline, message)
		case strings.EqualFold("slack", value):
			slack.Execute(headline, message)
		case strings.EqualFold("trello", value):
			trello.Execute(headline, message)
		case strings.EqualFold("zendesk", value):
			zendesk.Execute(headline, message)
		case strings.EqualFold("osticket", value):
			osticket.Execute(headline, message)
		case strings.EqualFold("twillo", value):
			twillo.Execute(headline, message)
		case strings.EqualFold("smtp", value):
			smtp.Execute(headline, message)
		case strings.EqualFold("webhook", value):
			webhook.Execute(headline, message)
		}
	}

}
