package webhook_test

import (
	"github.com/echgo/echgo/webhook"
	"os"
	"testing"
)

// TestExecute is to test the execute function
// We set the environment from the local variables
// And send a testing message to the service
func TestExecute(t *testing.T) {

	url := ""

	err := os.Setenv("WEBHOOK_URL", url)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	webhook.Execute(headline, message)

}
