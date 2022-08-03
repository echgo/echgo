package configuration

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

// Import is to import the configuration file
// An open, decode & close the json file
// Then we use reflect to create environments with regex
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

	regex := regexp.MustCompile("(.*?:)*(\")")

	channels := reflect.ValueOf(&Data.Channels).Elem()
	for i := 0; i < channels.NumField(); i++ {

		tag := string(channels.Type().Field(i).Tag)

		suffix := regex.ReplaceAllString(tag, "")
		suffix = strings.ToUpper(suffix)

		channel := reflect.ValueOf(channels.Field(i).Interface())
		for i := 0; i < channel.NumField(); i++ {

			tag := string(channel.Type().Field(i).Tag)

			prefix := regex.ReplaceAllString(tag, "")
			prefix = strings.ToUpper(prefix)

			content := fmt.Sprintf("%v", channel.Field(i).Interface())

			if len(content) > 0 {
				key := strings.Join([]string{suffix, prefix}, "_")
				err := os.Setenv(key, content)
				if err != nil {
					log.Fatalln(err)
				}
			}

		}

	}

}
