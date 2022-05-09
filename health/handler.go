package health

import (
	"log"
	"net/http"
	"os"
)

// paths are to define the directory & file paths for the stat lookup
var (
	paths = []string{"files/configuration", "files/configuration/echgo.yaml", "files/notification"}
)

// Handler is to serve the data from the health check
// First we define some variables for the return values
// After that, we check each file or directory & return a "good" or "bad"
func Handler(w http.ResponseWriter, r *http.Request) {

	health := true
	message := "bad"

	for _, value := range paths {
		if _, err := os.Stat(value); os.IsNotExist(err) {
			health = false
		}
	}

	if health {
		message = "good"
	}

	w.Header().Add("Content-Type", "text/plain")
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Fatalln(err)
	}

}
