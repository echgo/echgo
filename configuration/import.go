package configuration

import (
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

// Import is to notification the configuration
// To open, decode & close the yaml file
func Import() {

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

	err = yaml.Unmarshal(read, &Data)
	if err != nil {
		log.Fatalln(err)
	}

}
