package trello

import (
	"github.com/echgo/echgo/configuration"
	_ "github.com/echgo/echgo/testing"
	"github.com/echgo/echgo/trello"
	"testing"
)

// TestExecute is to test the trello execute function
func TestExecute(t *testing.T) {

	configuration.Import()

	headline := "Testing"
	message := "This is a message about test corners."

	trello.Execute(headline, message)

}
