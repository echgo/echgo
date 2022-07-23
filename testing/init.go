package testing

import (
	"log"
	"os"
	"path"
	"runtime"
)

// init is to get the right working directory for testing
func init() {

	_, filename, _, _ := runtime.Caller(0)

	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		log.Fatalln(err)
	}

}
