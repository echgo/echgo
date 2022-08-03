package gotify_test

import (
	"github.com/echgo/echgo/gotify"
	"os"
	"testing"
)

// TestExecute is to test the execute function
// We set the environment from the local variables
// And send a testing message to the service
func TestExecute(t *testing.T) {

	domain := ""
	key := ""

	err := os.Setenv("GOTIFY_DOMAIN", domain)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("GOTIFY_KEY", key)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	gotify.Execute(headline, message)

}
