package telegram

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/telegram"
	_ "github.com/echgo/echgo/testing"
	"testing"
)

// TestExecute is to test the telegram execute function
func TestExecute(t *testing.T) {

	configuration.Import()

	headline := "Testing"
	message := "This is a message about test corners."

	telegram.Execute(headline, message)

}
