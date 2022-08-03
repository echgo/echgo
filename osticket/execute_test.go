package osticket_test

import (
	"github.com/echgo/echgo/osticket"
	"os"
	"testing"
)

// baseUrl, apiToken, username, emailRecipient are important for the testing
// Please fill them out for successfully test
const (
	baseUrl        = ""
	apiToken       = ""
	username       = ""
	emailRecipient = ""
)

// TestExecute is to test the execute function
// We set the environment from the local const's
// And send a testing message to the service
func TestExecute(t *testing.T) {

	err := os.Setenv("OS_TICKET_BASE_URL", baseUrl)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("OS_TICKET_API_TOKEN", apiToken)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("OS_TICKET_USERNAME", username)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("OS_TICKET_EMAIL_RECIPIENT", emailRecipient)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	osticket.Execute(headline, message)

}
