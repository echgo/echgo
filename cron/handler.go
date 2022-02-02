package cron

import (
	"fmt"
	"github.com/adhocore/gronx"
	"github.com/echgo/echgo/configuration"
	"log"
)

// Handler is to handle the cronjobs & execute these
func Handler() {

	for _, value := range configuration.Data.Cronjobs {

		gron := gronx.New()

		due, err := gron.IsDue(value.Cron)
		if err != nil {
			log.Fatalln(err)
		}

		if due {
			fmt.Println(value.Name)
		}

	}

}
