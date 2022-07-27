package gotify

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/gotify"
	_ "github.com/echgo/echgo/testing"
	"testing"
)

// TestExecute is to test the gotify execute function
func TestExecute(t *testing.T) {

	configuration.Import()

	headline := "Testing"
	message := "This is a message about test corners."

	gotify.Execute(headline, message)

}
