package trello_test

import (
	"github.com/echgo/echgo/trello"
	"os"
	"testing"
)

// TestExecute is to test the execute function
// We set the environment from the local variables
// And send a testing message to the service
func TestExecute(t *testing.T) {

	webhookUrl := ""

	err := os.Setenv("SLACK_WEBHOOK_URL", webhookUrl)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	trello.Execute(headline, message)

}