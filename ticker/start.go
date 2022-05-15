package ticker

import (
	"time"
)

// Config is to structure the data
type Config struct {
	Time time.Time
}

// Start is to start the ticker with a job function every 15 seconds
func (c *Config) Start(job func()) {

	second := c.Time.Second()
	milli := c.Time.Nanosecond() / 1000000

	wait := 60*1000 - (second*1000 + milli)
	time.Sleep(time.Duration(wait) * time.Millisecond)

	for range time.Tick(time.Minute / 4) {
		go job()
	}

}
