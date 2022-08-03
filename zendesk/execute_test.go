package zendesk_test

import (
	"github.com/echgo/echgo/zendesk"
	"os"
	"testing"
)

// domain are important for the testing
// Please fill them out for successfully test
const (
	baseUrl  = ""
	username = ""
	apiToken = ""
)

// TestExecute is to test the execute function
// We set the environment from the local const's
// And send a testing message to the service
func TestExecute(t *testing.T) {

	err := os.Setenv("ZENDESK_BASE_URL", baseUrl)
	if err != nil {
		t.Log(err)
	}

	err = os.Setenv("ZENDESK_USERNAME", username)
	if err != nil {
		t.Log(err)
	}

	err = os.Setenv("ZENDESK_API_TOKEN", apiToken)
	if err != nil {
		t.Log(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	zendesk.Execute(headline, message)

}
