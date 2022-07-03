package configuration

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

// CreateIfNotExists is to create a configuration file, if none exists
// We add dummy data to the file
func CreateIfNotExists() {

	path := absolutePath()

	if _, err := os.Stat(path); os.IsNotExist(err) {

		data := Body{}

		prepare, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}

		err = os.WriteFile(path, prepare, 0666)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("A new configuration file has been created at /etc/echgo/configuration. Please fill in the configuration file and restart the container.")

		if flag.Lookup("test.v") == nil {
			os.Exit(1)
		}

	}

}
