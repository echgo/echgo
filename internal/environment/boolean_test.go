// Copyright 2024 Jonas Kwiedor. All rights reserved.

package environment_test

import (
	"github.com/echgo/echgo/v2/internal/environment"
	"testing"
)

// TestBoolean is to test the environment boolean function.
// We set the environment from the local variables & test
// the function. The return value is output in logging.
func TestBoolean(t *testing.T) {

	name := "USE_ENV"
	content := "true"

	t.Setenv(name, content)

	variable := environment.Boolean(name)

	t.Logf("The gotten value of the environment variable with the name \"%s\" is \"%v\". The expected value was \"%s\".\n", name, variable, content)

}
