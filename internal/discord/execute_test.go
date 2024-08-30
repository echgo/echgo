// Copyright 2024 Jonas Kwiedor. All rights reserved.

//go:build development

// Package discord_test is for testing the discord package.
package discord_test

import (
	"github.com/echgo/echgo/v2/internal/discord"
	"testing"
)

// TestExecute is to test the execute function. We set
// the environment from the local variables and
// send a testing message to the service.
func TestExecute(t *testing.T) {

	headline := "Testing"
	message := "This is a message about test corners."

	discord.Execute(headline, message)

}
