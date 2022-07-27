package configuration

import (
	"github.com/echgo/echgo/configuration"
	_ "github.com/echgo/echgo/testing"
	"testing"
)

// TestImport is to test the configuration import
func TestImport(t *testing.T) {
	configuration.Import()
}
