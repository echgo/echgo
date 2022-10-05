// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package zendesk_test

import (
	"github.com/echgo/echgo/v2/zendesk"
	"os"
	"testing"
)

// TestExecute is to test the execute function.
// We set the environment from the local variables
// and send a testing message to the service
func TestExecute(t *testing.T) {

	baseUrl := ""
	username := ""
	apiToken := ""

	err := os.Setenv("ZENDESK_BASE_URL", baseUrl)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("ZENDESK_USERNAME", username)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("ZENDESK_API_TOKEN", apiToken)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	zendesk.Execute(headline, message)

}
