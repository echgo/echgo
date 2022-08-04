package environment_test

import (
	"github.com/echgo/echgo/environment"
	"os"
	"testing"
)

// TestLookup is to test the environment lookup function
// We set the environment from the local variables & test the function
// The return value is output in logging
func TestLookup(t *testing.T) {

	name := "INTERVALL"
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
