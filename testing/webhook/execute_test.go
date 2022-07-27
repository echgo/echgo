package webhook

import (
	"github.com/echgo/echgo/configuration"
	_ "github.com/echgo/echgo/testing"
	"github.com/echgo/echgo/webhook"
	"testing"
)

// TestExecute is to test the webhook execute function
func TestExecute(t *testing.T) {

	configuration.Import()

	headline := "Testing"
	message := "This is a message about test corners."

	webhook.Execute(headline, message)

}
