package webhook_test

import (
	"github.com/echgo/echgo/webhook"
	"os"
	"testing"
)

// domain are important for the testing
// Please fill them out for successfully test
const (
	domain = ""
)

// TestExecute is to test the execute function
// We set the environment from the local const's
// And send a testing message to the service
func TestExecute(t *testing.T) {

	err := os.Setenv("WEBHOOK_DOMAIN", domain)
	if err != nil {
		t.Log(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	webhook.Execute(headline, message)

}
