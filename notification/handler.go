package notification

import (
	"github.com/echgo/echgo/configuration"
	"log"
	"os"
	"path/filepath"
)

// Handler is to handle the notification of file messaging files
// To open the file, check the file extension & remove the files
func Handler() {

	configuration.Import()

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	for _, value := range files {

		if value.Name() != ".gitkeep" && !value.IsDir() {

			log.Println("New notification received. These are now being processed.")

			file, err := os.Open(path + value.Name())
			if err != nil {
				log.Fatalln(err)
			}

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
