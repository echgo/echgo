package configuration

import (
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

// Import is to import the configuration
// To open & decode the yaml file
func Import() {

	open, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	read, err := io.ReadAll(open)
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(read, &Data)
	if err != nil {
		log.Fatalln(err)
	}

}
