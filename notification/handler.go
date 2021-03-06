package notification

import (
	"github.com/echgo/echgo/configuration"
	"github.com/echgo/echgo/console"
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
