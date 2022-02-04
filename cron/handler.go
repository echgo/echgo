package cron

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/notification"
	"log"
	"os/exec"
	"time"
)

// Handler is to handle the cronjobs & execute these
// Wo also import the configuration each hour
func Handler() {

	date := time.Now()

	configuration.Import()

	for index, value := range configuration.Data.Cronjobs {

		if IsDue(value.Cron, date) {

			cmd, err := exec.Command("/bin/bash", value.Path).Output()
			if err != nil {
				log.Fatalln(err)
			}

			response := string(cmd)

			if cmd != nil && len(response) > 0 {
				notification.Handler(response, index)
			}

		}

	}

}
