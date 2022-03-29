package notification

import (
	"bufio"
	"github.com/echgo/echgo/channels"
	"log"
	"os"
	"strings"
)

// TXT is to handle the txt files from the notification
// Check them & send them to the channel handler
// Them remove the files
func TXT(path string) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	parameter := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key, value, found := strings.Cut(scanner.Text(), "=")
		if found {
			parameter[key] = value
		}
	}

	if len(parameter["channel"]) > 0 {

		channel := channels.Type{}

		split := strings.Split(parameter["channel"], ",")
		for _, value := range split {
			trim := strings.TrimSpace(value)
			switch trim {
			case "gotify":
				channel.Gotify = true
			case "matrix":
				channel.Matrix = true
			case "telegram":
				channel.Telegram = true
			case "discord":
				channel.Discord = true
			case "slack":
				channel.Slack = true
			case "trello":
				channel.Trello = true
			case "zendesk":
				channel.Zendesk = true
			case "osticket":
				channel.Osticket = true
			case "smtp":
				channel.SMTP = true
			case "webhook":
				channel.Webhook = true
			}
		}

		channels.Handler(parameter["headline"], parameter["message"], channel)

	}

	err = os.Remove(path)
	if err != nil {
		log.Fatalln(err)
	}

}
