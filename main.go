package main

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/environment"
	"github.com/echgo/echgo/notification"
	"github.com/echgo/echgo/ticker"
	"time"
)

// Initialise the software with a console information
func init() {
	console.Init()
}

// If no env we add a dummy configuration file, if none exists
// And start ticker function with notification handler
func main() {

	if !environment.Boolean("USE_ENVIRONMENT") {
		configuration.Create()
	}

	c := ticker.Config{Time: time.Now()}
	c.Start(notification.Handler)

}
