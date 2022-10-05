// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package smtp_test

import (
	"github.com/echgo/echgo/v2/smtp"
	"os"
	"testing"
)

// TestExecute is to test the execute function.
// We set the environment from the local const's
// and send a testing message to the service.
func TestExecute(t *testing.T) {

	host := ""
	port := ""
	username := ""
	password := ""
	emailRecipient := ""

	err := os.Setenv("SMTP_HOST", host)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("SMTP_PORT", port)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("SMTP_USERNAME", username)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("SMTP_PASSWORD", password)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("SMTP_EMAIL_RECIPIENT", emailRecipient)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	smtp.Execute(headline, message)

}
