package gotify

import (
	"bytes"
	"fmt"
	"net/http"
)

// Config is to define config data
type Config struct {
	Path, Method string
	Body         []byte
}

// Request is to define the request data
type Request struct {
	Domain, XGotifyKey string
}

// Send is to send a new request
// & return the response
func (c *Config) Send(r Request) (*http.Response, error) {

	url := fmt.Sprintf("%s/%s", r.Domain, c.Path)

	client := &http.Client{}

	request, err := http.NewRequest(c.Method, url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Gotify-Key", r.XGotifyKey)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil

}
