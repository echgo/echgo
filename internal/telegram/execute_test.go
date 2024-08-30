// Copyright 2024 Jonas Kwiedor. All rights reserved.

//go:build development

// Package telegram_test is for testing the telegram package.
package telegram_test

import (
	"github.com/echgo/echgo/v2/internal/telegram"
	"testing"
)

// TestExecute is to test the execute function.
// We set the environment from the local variables
// and send a testing message to the service.
func TestExecute(t *testing.T) {

	headline := "Testing"
	message := "This is a message about test corners."

	telegram.Execute(headline, message)

}
