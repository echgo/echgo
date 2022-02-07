package notification

import (
	"github.com/echgo/echgo/configuration"
	"log"
	"os"
	"path/filepath"
)

// Handler is to handle the notification of file messaging files
func Handler() {

	configuration.Import()

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	for _, value := range files {

		if value.Name() != ".gitkeep" && !value.IsDir() {

			switch filepath.Ext(value.Name()) {
			case ".txt":
				TXT(path + value.Name())
			}

		}

	}

}
