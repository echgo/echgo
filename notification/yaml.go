package notification

import (
	"github.com/echgo/echgo/channels"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

// YamlBody is to decode the yaml file
type YamlBody struct {
	Channels []string `yaml:"channels"`
	Headline string   `yaml:"headline"`
	Message  string   `yaml:"message"`
}

// Yaml is to handle the yaml files from the notification
// Check them & send them to the channel handler
func Yaml(file *os.File) {

	var decode YamlBody

	read, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(read, &decode)
	if err != nil {
		log.Fatalln(err)
	}

	var types []string
	for _, value := range decode.Channels {
		types = append(types, value)
	}

	channels.Handler(decode.Headline, decode.Message, types)

}
