package smtp

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/smtp"
	_ "github.com/echgo/echgo/testing"
	"testing"
)

// TestExecute is to test the smtp execute function
func TestExecute(t *testing.T) {

	configuration.Import()

	headline := "Testing"
	message := "This is a message about test corners."

	smtp.Execute(headline, message)

}
