package main

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/notification"
	"github.com/echgo/echgo/ticker"
	"time"
)

// Initialise the software with a console information
func init() {
	console.Init()
}

// Add dummy configuration file, if none exists
// And start ticker as go function with notification handler
func main() {

	configuration.CreateIfNotExists()

	c := ticker.Config{Time: time.Now()}
	c.Start(notification.Handler)

}
