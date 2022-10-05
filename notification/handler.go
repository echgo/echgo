// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package notification

import (
	"github.com/echgo/echgo/v2/configuration"
	"github.com/echgo/echgo/v2/console"
	"github.com/echgo/echgo/v2/environment"
	"log"
	"os"
	"path/filepath"
)

// Handler is to handle the notification of file messaging files.
// To open the file, check the file extension & remove the files.
func Handler() {

	if !environment.Boolean("USE_ENVIRONMENT") {
		configuration.Import()
	}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	for _, value := range files {

		if !value.IsDir() {

			file, err := os.Open(path + value.Name())
			if err != nil {
				log.Fatalln(err)
			}

			attributes := make(map[string]any)
			attributes["file_name"] = value.Name()
			console.Log("info", "A file was imported.", attributes)

			switch filepath.Ext(value.Name()) {
			case ".txt":
				Txt(file)
			case ".json":
				Json(file)
			case ".xml":
				Xml(file)
			}

			err = file.Close()
			if err != nil {
				log.Fatalln(err)
			}

			err = os.Remove(path + value.Name())
			if err != nil {
				log.Fatalln(err)
			}

		}

	}

}
