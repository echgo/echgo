package configuration

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// CreateIfNotExists is to create a configuration file, if none exists
// We add dummy data to the file
func CreateIfNotExists() {

	if _, err := os.Stat(path); os.IsNotExist(err) {

		data := Body{}

		prepare, err := yaml.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}

		err = os.WriteFile(path, prepare, 0666)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("A new configuration file has been created at /etc/echgo/configuration. Please fill in the configuration file and restart the container.")

		os.Exit(0)

	}

}
