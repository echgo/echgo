package main

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/cron"
	"github.com/go-co-op/gocron"
	"time"
)

func init() {

	// Send init to console
	console.Init()

	// Import configuration
	configuration.Import()

}

func main() {

	// Define cronjobs &  start scheduler
	s := gocron.NewScheduler(time.UTC)

	s.Cron("* * * * *").Do(cron.Handler)

	s.StartBlocking()

}
