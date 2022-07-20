package configuration

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Import is to notification the configuration
// To open, decode & close the json file
func Import() {

	path := filepath.Join("files", "configuration", "default.json")

	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}

	read, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(read, &Data)
	if err != nil {
		log.Fatalln(err)
	}

}
