package main

import (
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/cron"
	"github.com/echgo/echgo/ticker"
	"time"
)

// Initialise the software with a console information
func init() {
	console.Init()
}

// Start ticker function with cron handler
func main() {
	c := ticker.Config{Time: time.Now()}
	c.Start(cron.Handler)
}
