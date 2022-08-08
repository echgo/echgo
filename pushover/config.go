package pushover

import (
	"bytes"
	"net/http"
)

// Set the base url for pushover service
const baseUrl = "https://api.pushover.net"

// Config is to define config data
type Config struct {
	Url, Method string
	Body        []byte
}

// Send is to send a new request
// & return the response
func (c *Config) Send() (*http.Response, error) {

	client := &http.Client{}

	request, err := http.NewRequest(c.Method, c.Url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil

}
