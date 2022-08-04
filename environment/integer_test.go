package environment_test

import (
	"github.com/echgo/echgo/environment"
	"os"
	"testing"
)

// TestInteger is to test the environment integer function
// We set the environment from the local variables & test the function
// The return value is output in logging
func TestInteger(t *testing.T) {

	name := "INTERVALL"
	content := "15"

	err := os.Setenv(name, content)
	if err != nil {
		t.Fatal(err)
	}

	variable := environment.Integer(name)

	t.Logf("The gotten value of the environment variable with the name \"%s\" is \"%d\". The expected value was \"%s\".\n", name, variable, content)

}
