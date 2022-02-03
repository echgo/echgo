package cron

import (
	"github.com/adhocore/gronx"
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/notification"
	"log"
	"os/exec"
)

// Handler is to handle the cronjobs & execute these
func Handler() {

	for index, value := range configuration.Data.Cronjobs {

		gron := gronx.New()

		due, err := gron.IsDue(value.Cron)
		if err != nil {
			log.Fatalln(err)
		}

		if due {

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
