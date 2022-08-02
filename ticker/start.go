package ticker

import (
	"github.com/echgo/echgo/environment"
	"time"
)

// Config is to structure the data
type Config struct {
	Time time.Time
}

// Start is to start the ticker with a job function every
// The default value is 15 to tick every 15 seconds
// You can overwrite this default with the environment `INTERVALL`
func (c *Config) Start(job func()) {

	second := c.Time.Second()
	milli := c.Time.Nanosecond() / 1000000

	wait := 60*1000 - (second*1000 + milli)
	time.Sleep(time.Duration(wait) * time.Millisecond)

	intervall := time.Duration(15)
	if environment.Integer("INTERVALL") > 0 {
		intervall = time.Duration(environment.Integer("INTERVALL"))
	}

	for range time.Tick(intervall * time.Second) {
		go job()
	}

}
