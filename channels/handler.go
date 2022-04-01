package channels

import (
	"github.com/echgo/echgo/discord"
	"github.com/echgo/echgo/gotify"
	"github.com/echgo/echgo/matrix"
	"github.com/echgo/echgo/osticket"
	"github.com/echgo/echgo/slack"
	"github.com/echgo/echgo/smtp"
	"github.com/echgo/echgo/telegram"
	"github.com/echgo/echgo/trello"
	"github.com/echgo/echgo/webhook"
	"github.com/echgo/echgo/zendesk"
	"strings"
)

// Handler is to handle the channels
// And send the notifications to them
func Handler(headline, message string, channel []string) {

	for _, value := range channel {
		switch {
		case strings.EqualFold("gotify", value):
			gotify.Execute(headline, message)
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
		case strings.EqualFold("smtp", value):
			smtp.Execute(headline, message)
		case strings.EqualFold("webhook", value):
			webhook.Execute(headline, message)
		}
	}

}
