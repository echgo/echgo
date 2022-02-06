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
			Cronjobs: []Cronjobs{},
		}

		data.Cronjobs = append(data.Cronjobs, Cronjobs{
			Path: "/var/lib/test-file.sh",
			Cron: "*/5 * * * *",
			Notification: Notification{
				Gotify:   true,
				Telegram: true,
				SMTP:     true,
				Webhook:  true,
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

		log.Println("A new configuration file has been created.")

	}

}
