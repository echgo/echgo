// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package ticker is used to start a function
// as goroutine in an interval to the full minute.
package ticker

import (
	"github.com/echgo/echgo/v2/environment"
	"time"
)

// Config is to structure the data.
type Config struct {
	Time time.Time
}

// Start is to start the ticker with a job function every.
// The default value is 15 to tick every 15 seconds. You
// can overwrite this default with the environment `INTERVAL`.
func (c *Config) Start(job func()) {

	second := c.Time.Second()
	milli := c.Time.Nanosecond() / 1000000

	wait := 60*1000 - (second*1000 + milli)
	time.Sleep(time.Duration(wait) * time.Millisecond)

	interval := time.Duration(15)
	if environment.Integer("INTERVAL") > 0 {
		interval = time.Duration(environment.Integer("INTERVAL"))
	}

	for range time.Tick(interval * time.Second) {
		go job()
	}

}
