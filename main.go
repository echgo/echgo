// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"github.com/echgo/echgo/v2/configuration"
	"github.com/echgo/echgo/v2/console"
	"github.com/echgo/echgo/v2/environment"
	"github.com/echgo/echgo/v2/notification"
	"github.com/echgo/echgo/v2/ticker"
	"time"
)

// Initialise the software with a console information.
func init() {
	console.Initialize()
}

// If no env we add a dummy configuration file, if none
// exists and start ticker function with notification handler.
func main() {

	if !environment.Boolean("USE_ENVIRONMENT") {
		configuration.Create()
	}

	c := ticker.Config{Time: time.Now()}
	c.Start(notification.Handler)

}
