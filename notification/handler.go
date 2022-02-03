package notification

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/email"
	"github.com/echgo/echgo/gotify"
	"github.com/echgo/echgo/telegram"
	"github.com/echgo/echgo/webhook"
	"log"
	"os"
)

var (
	gotifyRequest   = gotify.Request{Domain: os.Getenv("GOTIFY_DOMAIN"), XGotifyKey: os.Getenv("GOTIFY_KEY")}
	telegramRequest = telegram.Request{ApiToken: os.Getenv("TELEGRAM_API_TOKEN")}
	smtpData        = email.Smtp{Host: os.Getenv("SMTP_HOST"), Port: os.Getenv("SMTP_PORT"), Username: os.Getenv("SMTP_USERNAME"), Password: os.Getenv("SMTP_PASSWORD")}
	webhookRequest  = webhook.Request{Domain: os.Getenv("WEBHOOK_DOMAIN")}
)

func Handler(cmd []byte, index int) {

	switch {
	case configuration.Data.Cronjobs[index].Notification.Gotify:

		body := gotify.CreateMessageBody{
			Priority: 0,
			Title:    "echGo",
			Message:  string(cmd),
			Extras:   nil,
		}

		_, err := gotify.CreateMessage(body, gotifyRequest)
		if err != nil {
			log.Fatalln(err)
		}
	case configuration.Data.Cronjobs[index].Notification.Telegram:

		_, err := telegram.CreateMessage(string(cmd), "1998631548", "Markdown", telegramRequest)
		if err != nil {
			log.Fatalln(err)
		}

	case configuration.Data.Cronjobs[index].Notification.SMTP:

		d := email.Data{
			Email:   "jonas.kwiedor@jj-ideenschmiede.de",
			Subject: "echGo",
			File:    nil,
			HTML:    string(cmd),
		}

		err := email.Send(d, smtpData)
		if err != nil {
			log.Fatalln(err)
		}

	case configuration.Data.Cronjobs[index].Notification.Webhook:

		body := webhook.SendBody{
			Message: string(cmd),
		}

		err := webhook.Send(body, webhookRequest)
		if err != nil {
			log.Fatalln(err)
		}

	}

}
