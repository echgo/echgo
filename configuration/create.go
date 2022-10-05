// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package configuration

import (
	"encoding/json"
	"flag"
	"github.com/echgo/echgo/v2/console"
	"log"
	"os"
	"path/filepath"
)

// Create is to create a configuration file, if
// none exists. We add dummy data to the file.
func Create() {

	path := filepath.Join("files", "configuration", "default.json")

	if _, err := os.Stat(path); os.IsNotExist(err) {

		prepare, err := json.MarshalIndent(Body{}, "", "	")
		if err != nil {
			log.Fatalln(err)
		}

		err = os.WriteFile(path, prepare, 0666)
		if err != nil {
			log.Fatalln(err)
		}

		console.Log("info", "A new configuration file has been created at /etc/echgo/configuration.", map[string]any{})

		if flag.Lookup("test.v") == nil {
			os.Exit(1)
		}

	}

}
