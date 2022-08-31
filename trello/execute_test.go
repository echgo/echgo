package trello_test

import (
	"github.com/echgo/echgo/v2/trello"
	"os"
	"testing"
)

// TestExecute is to test the execute function
// We set the environment from the local variables
// And send a testing message to the service
func TestExecute(t *testing.T) {

	key := "b4c3c561bc8cf6c409fc385633c1b120"
	token := "d7b22153825871e8ad8001569e39dbd7d62f1cb1f18e9d238f3013efd7d43080"
	idList := "62e298ceed5e0d7c89d074ca"

	err := os.Setenv("TRELLO_KEY", key)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("TRELLO_TOKEN", token)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("TRELLO_ID_LIST", idList)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Update erfolgreich (creeb-demo)"
	message := "**Kunde**: \n**Lösung**: \n**Alte Version**: \n**Neue Version**:"

	trello.Execute(headline, message)

}
