// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package telegram_test

import (
	"github.com/echgo/echgo/v2/telegram"
	"os"
	"testing"
)

// TestExecute is to test the execute function.
// We set the environment from the local variables
// and send a testing message to the service.
func TestExecute(t *testing.T) {

	apiToken := ""
	chatId := ""

	err := os.Setenv("TELEGRAM_API_TOKEN", apiToken)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Setenv("TELEGRAM_CHAT_ID", chatId)
	if err != nil {
		t.Fatal(err)
	}

	headline := "Testing"
	message := "This is a message about test corners."

	telegram.Execute(headline, message)

}
