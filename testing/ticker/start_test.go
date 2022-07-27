package ticker

import (
	_ "github.com/echgo/echgo/testing"
	"github.com/echgo/echgo/ticker"
	"log"
	"testing"
	"time"
)

// TestStart is to test the ticker function
func TestStart(t *testing.T) {

	c := ticker.Config{Time: time.Now()}
	c.Start(func() {
		log.Println("The ticker was executed.")
	})

}
