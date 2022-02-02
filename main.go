package main

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/cron"
	"time"
)

func init() {

	// Send init to console
	console.Init()

	// Import configuration
	configuration.Import()

}

func main() {

	// Start ticker function with cron handler
	c := cron.Config{Time: time.Now()}
	c.Ticker(cron.Handler)

}
