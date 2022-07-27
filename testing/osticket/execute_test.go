package osticket

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/osticket"
	_ "github.com/echgo/echgo/testing"
	"testing"
)

// TestExecute is to test the osticket execute function
func TestExecute(t *testing.T) {

	configuration.Import()

	headline := "Testing"
	message := "This is a message about test corners."

	osticket.Execute(headline, message)

}
