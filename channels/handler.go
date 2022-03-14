package channels

import (
	"fmt"
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/discord"
	"github.com/echgo/echgo/email"
	"github.com/echgo/echgo/gotify"
	"github.com/echgo/echgo/matrix"
	"github.com/echgo/echgo/telegram"
	"github.com/echgo/echgo/trello"
	"github.com/echgo/echgo/webhook"
	"github.com/echgo/echgo/zendesk"
	"log"
)

// Handler is to handle the channels
// And send the notifications to them
func Handler(headline, message string, channel Type) {

	if channel.Gotify {

		r := gotify.Request{
			Domain:     configuration.Data.Channels.Gotify.Domain,
			XGotifyKey: configuration.Data.Channels.Gotify.Key,
		}

		b := gotify.CreateMessageBody{
			Priority: 0,
			Title:    headline,
			Message:  message,
			Extras:   nil,
		}

		_, err := gotify.CreateMessage(b, r)
		if err != nil {
			log.Fatalln(err)
		}

	}

	if channel.Matrix {

		r := matrix.Request{
			Domain:      configuration.Data.Channels.Matrix.Domain,
			RoomId:      configuration.Data.Channels.Matrix.RoomId,
			AccessToken: configuration.Data.Channels.Matrix.AccessToken,
		}

		b := matrix.CreateMessageBody{
			Msgtype: "m.text",
			Body:    fmt.Sprintf("%s\n%s", headline, message),
		}

		_, err := matrix.CreateMessage(b, r)
		if err != nil {
			log.Fatalln(err)
		}

	}

	if channel.Telegram {

		r := telegram.Request{
			ApiToken: configuration.Data.Channels.Telegram.ApiToken,
			ChatId:   configuration.Data.Channels.Telegram.ChatId,
		}

		_, err := telegram.CreateMessage(fmt.Sprintf("%s\n%s", headline, message), "Markdown", r)
		if err != nil {
			log.Fatalln(err)
		}

	}

	if channel.Discord {

		r := discord.Request{
			Url: configuration.Data.Channels.Discord.WebhookUrl,
		}

		if len(configuration.Data.Channels.Discord.BotName) == 0 {
			configuration.Data.Channels.Discord.BotName = "echGo"
		}

		b := discord.CreateMessageBody{
			Username:  configuration.Data.Channels.Discord.BotName,
			AvatarUrl: "https://raw.githubusercontent.com/echgo/logo/main/echGo-small.png",
			Content:   fmt.Sprintf("%s - %s", headline, message),
		}

		err := discord.CreateMessage(b, r)
		if err != nil {
			log.Fatalln(err)
		}

	}

	if channel.Trello {

		r := trello.Request{
			Key:    configuration.Data.Channels.Trello.Key,
			Token:  configuration.Data.Channels.Trello.Token,
			IdList: configuration.Data.Channels.Trello.IdList,
		}

		_, err := trello.CreateCard(headline, message, r)
		if err != nil {
			log.Fatalln(err)
		}

	}

	if channel.Zendesk {

		r := zendesk.Request{
			BaseUrl:  configuration.Data.Channels.Zendesk.BaseUrl,
			Username: configuration.Data.Channels.Zendesk.Username,
			ApiToken: configuration.Data.Channels.Zendesk.ApiToken,
		}

		b := zendesk.CreateTicketBody{
			Ticket: zendesk.CreateTicketBodyTicket{
				Comment: zendesk.CreateTicketBodyComment{
					Body: message,
				},
				Priority: "urgent",
				Subject:  headline,
			},
		}

		_, err := zendesk.CreateTicket(b, r)
		if err != nil {
			log.Fatalln(err)
		}

	}

	if channel.SMTP {

		s := email.Smtp{
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

		err := email.Send(d, s)
		if err != nil {
			log.Fatalln(err)
		}

	}

	if channel.Webhook {

		r := webhook.Request{
			Domain: configuration.Data.Channels.Webhook.Domain,
		}

		b := webhook.SendBody{
			Headline: headline,
			Message:  message,
		}

		err := webhook.Send(b, r)
		if err != nil {
			log.Fatalln(err)
		}

	}

}
