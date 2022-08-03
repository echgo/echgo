package twillo_test

import (
	"github.com/echgo/echgo/twillo"
	"os"
	"testing"
)

// accountSid, authToken, phoneNumber, myPhoneNumber are important for the testing
// Please fill them out for successfully test
const (
	accountSid    = ""
	authToken     = ""
	phoneNumber   = ""
	myPhoneNumber = ""
)

// TestExecute is to test the execute function
// We set the environment from the local const's
// And send a testing message to the service
func TestExecute(t *testing.T) {

	err := os.Setenv("TWILLO_ACCOUNT_SID", accountSid)
	if err != nil {
		t.Log(err)
	}

	err = os.Setenv("TWILLO_AUTH_TOKEN", authToken)
	if err != nil {
		t.Log(err)
	}

	err = os.Setenv("TWILLO_PHONE_NUMBER", phoneNumber)
	if err != nil {
		t.Log(err)
	}

	err = os.Setenv("TWILLO_MY_PHONE_NUMBER", myPhoneNumber)
	if err != nil {
		t.Log(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	twillo.Execute(headline, message)

}
