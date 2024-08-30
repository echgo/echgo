// Copyright 2024 Jonas Kwiedor. All rights reserved.

// Package configuration_test is for testing the configuration package.
package configuration_test

import (
	"github.com/echgo/echgo/v2/internal/configuration"
	"log"
	"os"
	"path"
	"runtime"
	"testing"
)

// TestCreate is to test the creation of the configuration file.
func TestCreate(t *testing.T) {

	_, filename, _, _ := runtime.Caller(0)

	dir := path.Join(path.Dir(filename), "..", "..")
	err := os.Chdir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	configuration.Create()

}
