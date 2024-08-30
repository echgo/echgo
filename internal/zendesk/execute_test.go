// Copyright 2024 Jonas Kwiedor. All rights reserved.

//go:build development

// Package zendesk_test is for testing the zendesk package.
package zendesk_test

import (
	"github.com/echgo/echgo/v2/internal/zendesk"
	"testing"
)

// TestExecute is to test the execute function.
// We set the environment from the local variables
// and send a testing message to the service
func TestExecute(t *testing.T) {

	headline := "Testing"
	message := "This is a message about test corners."

	zendesk.Execute(headline, message)

}
