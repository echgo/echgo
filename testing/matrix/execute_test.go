package matrix

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/matrix"
	_ "github.com/echgo/echgo/testing"
	"testing"
)

// TestExecute is to test the matrix execute function
func TestExecute(t *testing.T) {

	configuration.Import()

	headline := "Testing"
	message := "This is a message about test corners."

	matrix.Execute(headline, message)

}
