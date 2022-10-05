// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package trello_test

import (
	"github.com/echgo/echgo/v2/trello"
	"os"
	"testing"
)

// TestExecute is to test the execute function.
// We set the environment from the local variables
// and send a testing message to the service
func TestExecute(t *testing.T) {

	key := ""
	token := ""
	idList := ""

	err := os.Setenv("TRELLO_KEY", key)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("TRELLO_TOKEN", token)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("TRELLO_ID_LIST", idList)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	trello.Execute(headline, message)

}
