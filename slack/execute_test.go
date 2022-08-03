package slack_test

import (
	"github.com/echgo/echgo/slack"
	"os"
	"testing"
)

// webhookUrl are important for the testing
// Please fill them out for successfully test
const (
	webhookUrl = ""
)

// TestExecute is to test the execute function
// We set the environment from the local const's
// And send a testing message to the service
func TestExecute(t *testing.T) {

	err := os.Setenv("SLACK_WEBHOOK_URL", webhookUrl)
	if err != nil {
		t.Log(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	slack.Execute(headline, message)

}
