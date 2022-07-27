package slack

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/slack"
	_ "github.com/echgo/echgo/testing"
	"testing"
)

// TestExecute is to test the slack execute function
func TestExecute(t *testing.T) {

	configuration.Import()

	headline := "Testing"
	message := "This is a message about test corners."

	slack.Execute(headline, message)

}
