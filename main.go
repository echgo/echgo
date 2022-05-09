package main

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
	"github.com/echgo/echgo/health"
	"github.com/echgo/echgo/notification"
	"github.com/echgo/echgo/ticker"
	"log"
	"net/http"
	"time"
)

// Initialise the software with a console information
func init() {
	console.Init()
}

// Add dummy configuration file, if none exists
// And start ticker as go function with notification handler
// Now start a little webserver for health checks
func main() {

	configuration.CreateIfNotExists()

	c := ticker.Config{Time: time.Now()}
	go c.Start(notification.Handler)

	http.HandleFunc("/health", health.Handler)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
