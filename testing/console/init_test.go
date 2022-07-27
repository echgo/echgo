package console

import (
	"github.com/echgo/echgo/console"
	_ "github.com/echgo/echgo/testing"
	"testing"
)

// TestInit is to test the console initialize function
func TestInit(t *testing.T) {
	console.Init()
}
