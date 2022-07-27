package zendesk

import (
	"github.com/echgo/echgo/configuration"
	_ "github.com/echgo/echgo/testing"
	"github.com/echgo/echgo/zendesk"
	"testing"
)

// TestExecute is to test the zendesk execute function
func TestExecute(t *testing.T) {

	configuration.Import()

	headline := "Testing"
	message := "This is a message about test corners."

	zendesk.Execute(headline, message)

}
