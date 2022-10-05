// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pushover_test

import (
	"github.com/echgo/echgo/v2/pushover"
	"os"
	"testing"
)

// TestExecute is to test the execute function.
// We set the environment from the local variables
// and send a testing message to the service
func TestExecute(t *testing.T) {

	token := ""
	user := ""

	err := os.Setenv("PUSHOVER_TOKEN", token)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("PUSHOVER_USER", user)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	pushover.Execute(headline, message)

}
