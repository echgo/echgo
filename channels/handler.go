package channels

import (
	"fmt"
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/email"
	"github.com/echgo/echgo/gotify"
	"github.com/echgo/echgo/telegram"
	"github.com/echgo/echgo/webhook"
	"log"
)

// Handler is to handle the channels
// And send the notifications to them
func Handler(headline, message string, channel Type) {

	if channel.Gotify {

		gotifyRequest := gotify.Request{
			Domain:     configuration.Data.Channels.Gotify.Domain,
			XGotifyKey: configuration.Data.Channels.Gotify.Key,
		}

		body := gotify.CreateMessageBody{
			Priority: 0,
			Title:    headline,
			Message:  message,
			Extras:   nil,
		}

		_, err := gotify.CreateMessage(body, gotifyRequest)
		if err != nil {
			log.Fatalln(err)
		}

	}

	if channel.Telegram {

		telegramRequest := telegram.Request{
			ApiToken: configuration.Data.Channels.Telegram.ApiToken,
		}

		_, err := telegram.CreateMessage(fmt.Sprintf("%s\n%s", headline, message), configuration.Data.Channels.Telegram.ChatId, "Markdown", telegramRequest)
		if err != nil {
			log.Fatalln(err)
		}

	}

	if channel.SMTP {

		smtpData := email.Smtp{
			Host:     configuration.Data.Channels.SMTP.Host,
			Port:     configuration.Data.Channels.SMTP.Port,
			Username: configuration.Data.Channels.SMTP.Username,
			Password: configuration.Data.Channels.SMTP.Password,
		}

		d := email.Data{
			Email:   configuration.Data.Channels.SMTP.EmailRecipient,
			Subject: headline,
			File:    nil,
			HTML:    message,
		}

		err := email.Send(d, smtpData)
		if err != nil {
			log.Fatalln(err)
		}

	}

	if channel.Webhook {

		webhookRequest := webhook.Request{
			Domain: configuration.Data.Channels.Webhook.Domain,
		}

		body := webhook.SendBody{
			Headline: headline,
			Message:  message,
		}

		err := webhook.Send(body, webhookRequest)
		if err != nil {
			log.Fatalln(err)
		}

	}

}
