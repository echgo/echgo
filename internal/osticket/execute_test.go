// Copyright 2024 Jonas Kwiedor. All rights reserved.

//go:build development

// Package osticket_test is for testing the osticket package.
package osticket_test

import (
	"github.com/echgo/echgo/v2/internal/osticket"
	"testing"
)

// TestExecute is to test the execute function.
// We set the environment from the local variables
// and send a testing message to the service
func TestExecute(t *testing.T) {

	headline := "Testing"
	message := "This is a message about test corners."

	osticket.Execute(headline, message)

}
