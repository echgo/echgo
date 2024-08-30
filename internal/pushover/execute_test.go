// Copyright 2024 Jonas Kwiedor. All rights reserved.

//go:build development

// Package pushover_test is for testing the pushover package.
package pushover_test

import (
	"github.com/echgo/echgo/v2/internal/pushover"
	"testing"
)

// TestExecute is to test the execute function.
// We set the environment from the local variables
// and send a testing message to the service
func TestExecute(t *testing.T) {

	headline := "Testing"
	message := "This is a message about test corners."

	pushover.Execute(headline, message)

}
