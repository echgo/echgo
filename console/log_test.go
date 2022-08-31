package console_test

import (
	"github.com/echgo/echgo/v2/console"
	"testing"
)

// TestLog is to test the console log function
func TestLog(t *testing.T) {

	attributes := make(map[string]any)
	attributes["file_name"] = "test-file.txt"
	console.Log("info", "A file was imported.", attributes)

}
