package channels

import (
	"github.com/echgo/echgo/discord"
	"github.com/echgo/echgo/email"
	"github.com/echgo/echgo/gotify"
	"github.com/echgo/echgo/matrix"
	"github.com/echgo/echgo/osticket"
	"github.com/echgo/echgo/slack"
	"github.com/echgo/echgo/telegram"
	"github.com/echgo/echgo/trello"
	"github.com/echgo/echgo/webhook"
	"github.com/echgo/echgo/zendesk"
)

// Handler is to handle the channels
// And send the notifications to them
func Handler(headline, message string, channel Type) {

	if channel.Gotify {
		gotify.Execute(headline, message)
	}

	if channel.Matrix {
		matrix.Execute(headline, message)
	}

	if channel.Telegram {
		telegram.Execute(headline, message)
	}

	if channel.Discord {
		discord.Execute(headline, message)
	}

	if channel.Slack {
		slack.Execute(headline, message)
	}

	if channel.Trello {
		trello.Execute(headline, message)
	}

	if channel.Zendesk {
		zendesk.Execute(headline, message)
	}

	if channel.Osticket {
		osticket.Execute(headline, message)
	}

	if channel.SMTP {
		email.Execute(headline, message)
	}

	if channel.Webhook {
		webhook.Execute(headline, message)
	}

}
