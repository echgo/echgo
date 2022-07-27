package configuration

import (
	"github.com/echgo/echgo/configuration"
	_ "github.com/echgo/echgo/testing"
	"testing"
)

// TestCreate is to test the creation of the configuration file
func TestCreate(t *testing.T) {
	configuration.Create()
}
