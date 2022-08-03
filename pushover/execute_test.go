package pushover_test

import (
	"github.com/echgo/echgo/pushover"
	"os"
	"testing"
)

// token, user are important for the testing
// Please fill them out for successfully test
const (
	token = ""
	user  = ""
)

// TestExecute is to test the execute function
// We set the environment from the local const's
// And send a testing message to the service
func TestExecute(t *testing.T) {

	err := os.Setenv("PUSHOVER_TOKEN", token)
	if err != nil {
		t.Log(err)
	}

	err = os.Setenv("PUSHOVER_USER", user)
	if err != nil {
		t.Log(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	pushover.Execute(headline, message)

}
