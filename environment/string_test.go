// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package environment_test

import (
	"github.com/echgo/echgo/v2/environment"
	"os"
	"testing"
)

// TestString is to test the environment string function.
// We set the environment from the local variables & test
// the function. The return value is output in logging.
func TestString(t *testing.T) {

	name := "BOT_NAME"
	content := "echGo"

	err := os.Setenv(name, content)
	if err != nil {
		t.Fatal(err)
	}

	variable := environment.String(name)

	t.Logf("The gotten value of the environment variable with the name \"%s\" is \"%s\". The expected value was \"%s\".\n", name, variable, content)

}
