package notification

import (
	"encoding/xml"
	"github.com/echgo/echgo/channels"
	"io"
	"log"
	"os"
)

// XmlBody is to decode the xml file
type XmlBody struct {
	Channels struct {
		Type []string `xml:"type"`
	} `xml:"channels"`
	Headline string `xml:"headline"`
	Message  string `xml:"message"`
}

// Xml is to handle the xml files from the notification
// Check them & send them to the channel handler
func Xml(file *os.File) {

	var decode XmlBody

	read, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	err = xml.Unmarshal(read, &decode)
	if err != nil {
		log.Fatalln(err)
	}

	var types []string
	types = append(types, decode.Channels.Type...)

	channels.Handler(decode.Headline, decode.Message, types)

}
