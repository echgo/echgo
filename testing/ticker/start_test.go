package ticker

import (
	_ "github.com/echgo/echgo/testing"
	"github.com/echgo/echgo/ticker"
	"log"
	"testing"
	"time"
)

// TestStart is to test the ticker function
// Use an intervall to end the testing function `go test --timeout 280s`
func TestStart(t *testing.T) {

	t.Setenv("INTERVALL", "5")

	c := ticker.Config{Time: time.Now()}
	c.Start(func() {
		log.Println("The ticker was executed.")
	})

}
