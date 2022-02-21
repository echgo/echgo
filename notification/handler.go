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

			log.Println("New notification received. This will now be processed further.")

			switch filepath.Ext(value.Name()) {
			case ".txt":
				TXT(path + value.Name())
			case ".json":
				JSON(path + value.Name())
			}

		}

	}

}
