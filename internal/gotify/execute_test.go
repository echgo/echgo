// Copyright 2024 Jonas Kwiedor. All rights reserved.

//go:build development

// Package gotify_test is for testing the gotify package.
package gotify_test

import (
	"github.com/echgo/echgo/v2/internal/gotify"
	"testing"
)

// TestExecute is to test the execute function.
// We set the environment from the local variables
// and send a testing message to the service.
func TestExecute(t *testing.T) {

	headline := "Testing"
	message := "This is a message about test corners."

	gotify.Execute(headline, message)

}
