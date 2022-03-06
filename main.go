package main

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/notification"
	"github.com/echgo/echgo/ticker"
	"time"
)

// Initialise the software with a console information
// And add dummy configuration file, if none exists
func init() {
	console.Init()
	configuration.CreateIfNotExists()
}

// Start ticker function with notification handler
func main() {
	c := ticker.Config{Time: time.Now()}
	c.Start(notification.Handler)
}
