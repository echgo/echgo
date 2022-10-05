// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package configuration_test

import (
	"fmt"
	"github.com/echgo/echgo/v2/configuration"
	"log"
	"os"
	"path"
	"runtime"
	"testing"
)

// TestImport is to test the configuration import.
func TestImport(t *testing.T) {

	_, file, _, _ := runtime.Caller(0)

	dir := path.Join(path.Dir(file), "..")
	err := os.Chdir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	configuration.Import()

	fmt.Println(os.Environ())

}
