package pushover

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/pushover"
	_ "github.com/echgo/echgo/testing"
	"testing"
)

// TestExecute is to test the pushover execute function
func TestExecute(t *testing.T) {

	configuration.Import()

	headline := "Testing"
	message := "This is a message about test corners."

	pushover.Execute(headline, message)

}
