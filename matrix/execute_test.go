package matrix_test

import (
	"github.com/echgo/echgo/matrix"
	"os"
	"testing"
)

// TestExecute is to test the execute function
// We set the environment from the local variables
// And send a testing message to the service
func TestExecute(t *testing.T) {

	baseUrl := ""
	roomId := ""
	accessToken := ""

	err := os.Setenv("MATRIX_BASE_URL", baseUrl)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("MATRIX_ROOM_ID", roomId)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("MATRIX_ACCESS_TOKEN", accessToken)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	matrix.Execute(headline, message)

}
