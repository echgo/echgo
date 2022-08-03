package telegram_test

import (
	"github.com/echgo/echgo/telegram"
	"os"
	"testing"
)

// apiToken, chatId are important for the testing
// Please fill them out for successfully test
const (
	apiToken = ""
	chatId   = ""
)

// TestExecute is to test the execute function
// We set the environment from the local const's
// And send a testing message to the service
func TestExecute(t *testing.T) {

	err := os.Setenv("TELEGRAM_API_TOKEN", apiToken)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("TELEGRAM_CHAT_ID", chatId)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	telegram.Execute(headline, message)

}
