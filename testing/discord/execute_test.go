package discord

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/discord"
	_ "github.com/echgo/echgo/testing"
	"testing"
)

// TestExecute is to test the discord execute function
func TestExecute(t *testing.T) {

	configuration.Import()

	headline := "Testing"
	message := "This is a message about test corners."

	discord.Execute(headline, message)

}
