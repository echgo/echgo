// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

//go:build development

// Package twillo_test is for testing the twillo package.
package twillo_test

import (
	"github.com/echgo/echgo/v2/internal/twillo"
	"testing"
)

// TestExecute is to test the execute function.
// We set the environment from the local const's
// and send a testing message to the service.
func TestExecute(t *testing.T) {

	headline := "Testing"
	message := "This is a message about test corners."

	twillo.Execute(headline, message)

}
