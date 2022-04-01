package notification

import (
	"bufio"
	"github.com/echgo/echgo/channels"
	"os"
	"strings"
)

// Txt is to handle the txt files from the notification
// Check them & send them to the channel handler
func Txt(file *os.File) {

	parameter := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key, value, found := strings.Cut(scanner.Text(), "=")
		if found {
			parameter[key] = value
		}
	}

	if len(parameter["channel"]) > 0 {

		var channel []string

		split := strings.Split(parameter["channel"], ",")
		for _, value := range split {
			trim := strings.TrimSpace(value)
			channel = append(channel, trim)
		}

		channels.Handler(parameter["headline"], parameter["message"], channel)

	}

}
