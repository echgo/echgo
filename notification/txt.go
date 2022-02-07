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
		if strings.Contains(scanner.Text(), "=") {
			split := strings.Split(scanner.Text(), "=")
			parameter[split[0]] = split[1]
		}
	}

	if len(parameter["channel"]) > 0 && strings.Contains(parameter["channel"], ",") {
		channel := channels.Type{}
		split := strings.Split(parameter["channel"], ",")
		for _, value := range split {
			switch value {
			case "gotify":
				channel.Gotify = true
			case "telegram":
				channel.Telegram = true
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
