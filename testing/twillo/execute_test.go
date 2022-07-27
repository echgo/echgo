package twillo

import (
	"github.com/echgo/echgo/configuration"
	_ "github.com/echgo/echgo/testing"
	"github.com/echgo/echgo/twillo"
	"testing"
)

// TestExecute is to test the twillo execute function
func TestExecute(t *testing.T) {

	configuration.Import()

	headline := "Testing"
	message := "This is a message about test corners."

	twillo.Execute(headline, message)

}
