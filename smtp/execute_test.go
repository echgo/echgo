package smtp_test

import (
	"github.com/echgo/echgo/smtp"
	"os"
	"testing"
)

// host, port, username, password, emailRecipient are important for the testing
// Please fill them out for successfully test
const (
	host           = ""
	port           = ""
	username       = ""
	password       = ""
	emailRecipient = ""
)

// TestExecute is to test the execute function
// We set the environment from the local const's
// And send a testing message to the service
func TestExecute(t *testing.T) {

	err := os.Setenv("SMTP_HOST", host)
	if err != nil {
		t.Log(err)
	}

	err = os.Setenv("SMTP_PORT", port)
	if err != nil {
		t.Log(err)
	}

	err = os.Setenv("SMTP_USERNAME", username)
	if err != nil {
		t.Log(err)
	}

	err = os.Setenv("SMTP_PASSWORD", password)
	if err != nil {
		t.Log(err)
	}

	err = os.Setenv("SMTP_EMAIL_RECIPIENT", emailRecipient)
	if err != nil {
		t.Log(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	smtp.Execute(headline, message)

}
