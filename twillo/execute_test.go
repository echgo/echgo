package twillo_test

import (
	"github.com/echgo/echgo/v2/twillo"
	"os"
	"testing"
)

// TestExecute is to test the execute function
// We set the environment from the local const's
// And send a testing message to the service
func TestExecute(t *testing.T) {

	accountSid := ""
	authToken := ""
	phoneNumber := ""
	myPhoneNumber := ""

	err := os.Setenv("TWILLO_ACCOUNT_SID", accountSid)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("TWILLO_AUTH_TOKEN", authToken)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("TWILLO_PHONE_NUMBER", phoneNumber)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("TWILLO_MY_PHONE_NUMBER", myPhoneNumber)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	twillo.Execute(headline, message)

}
