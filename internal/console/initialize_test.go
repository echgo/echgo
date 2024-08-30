// Copyright 2024 Jonas Kwiedor. All rights reserved.

// Package console_test is for testing the console package.
package console_test

import (
	"github.com/echgo/echgo/v2/internal/console"
	"testing"
)

// TestInitialize is to test the console initialize function.
func TestInitialize(t *testing.T) {
	console.Initialize()
}
