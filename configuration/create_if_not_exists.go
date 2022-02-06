package configuration

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// CreateIfNotExists is to create a configuration file, if none exists
// We add dummy data to the file
func CreateIfNotExists() {

	if _, err := os.Stat(filePath); os.IsNotExist(err) {

		data := Body{
			NotificationChannels: NotificationChannels{
				Gotify:   Gotify{},
				Telegram: Telegram{},
			},
			Cronjobs: []Cronjobs{},
		}

		data.Cronjobs = append(data.Cronjobs, Cronjobs{
			Path: "/var/lib/test-file.sh",
			Cron: "0 8 * * 0",
			Notification: Notification{
				Gotify:   false,
				Telegram: false,
				SMTP:     false,
				Webhook:  false,
			},
		})

		prepare, err := yaml.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}

		err = os.WriteFile(filePath, prepare, 0666)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("A new configuration file has been created. Please fill in the configuration file and restart.")

		os.Exit(0)

	}

}
