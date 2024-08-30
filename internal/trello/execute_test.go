// Copyright 2024 Jonas Kwiedor. All rights reserved.

//go:build development

// Package trello_test is for testing the trello package.
package trello_test

import (
	"github.com/echgo/echgo/v2/internal/trello"
	"testing"
)

// TestExecute is to test the execute function.
// We set the environment from the local variables
// and send a testing message to the service
func TestExecute(t *testing.T) {

	headline := "Testing"
	message := "This is a message about test corners."

	trello.Execute(headline, message)

}
