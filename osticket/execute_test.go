// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package osticket_test

import (
	"github.com/echgo/echgo/v2/osticket"
	"os"
	"testing"
)

// TestExecute is to test the execute function.
// We set the environment from the local variables
// and send a testing message to the service
func TestExecute(t *testing.T) {

	baseUrl := ""
	apiToken := ""
	username := ""
	emailRecipient := ""

	err := os.Setenv("OS_TICKET_BASE_URL", baseUrl)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("OS_TICKET_API_TOKEN", apiToken)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("OS_TICKET_USERNAME", username)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("OS_TICKET_EMAIL_RECIPIENT", emailRecipient)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	osticket.Execute(headline, message)

}
