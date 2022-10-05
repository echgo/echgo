// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package environment_test

import (
	"github.com/echgo/echgo/v2/environment"
	"os"
	"testing"
)

// TestLookup is to test the environment lookup function.
// We set the environment from the local variables & test
// the function. The return value is output in logging.
func TestLookup(t *testing.T) {

	name := "INTERVAL"
	content := "15"

	err := os.Setenv(name, content)
	if err != nil {
		t.Fatal(err)
	}

	lookup := environment.Lookup(name)
	if lookup {
		t.Logf("The environment variable with the name \"%s\" exists.\n", name)
	}

}
