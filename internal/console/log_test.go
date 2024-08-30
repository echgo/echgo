// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package console_test

import (
	"github.com/echgo/echgo/v2/internal/console"
	"testing"
)

// TestLog is to test the console log function.
func TestLog(t *testing.T) {

	attributes := make(map[string]any)
	attributes["file_name"] = "test-file.txt"
	console.Log("info", "A file was imported.", attributes)

}
