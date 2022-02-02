package cron

import (
	"time"
)

// Config is to structure the data
type Config struct {
	Time time.Time
}

// Ticker is to start a job function every minute
func (c *Config) Ticker(job func()) {

	second := c.Time.Second()
	nano := c.Time.Nanosecond()

	wait := 60*1000000000 - (second*1000000000 + nano)
	time.Sleep(time.Duration(wait) * time.Nanosecond)

	for _ = range time.Tick(time.Minute) {
		job()
	}

}
