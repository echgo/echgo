package discord_test

import (
	"github.com/echgo/echgo/discord"
	"os"
	"testing"
)

// TestExecute is to test the execute function
// We set the environment from the local variables
// And send a testing message to the service
func TestExecute(t *testing.T) {

	webhookUrl := ""
	botName := ""

	err := os.Setenv("DISCORD_WEBHOOK_URL", webhookUrl)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("DISCORD_BOT_NAME", botName)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	discord.Execute(headline, message)

}