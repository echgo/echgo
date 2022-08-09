package gotify

import (
	"bytes"
	"net/http"
)

// channel is to save the channel name for logging
const channel = "gotify"

// Config is to define config data
type Config struct {
	Url, Method string
	Body        []byte
}

// Request is to define the request data
type Request struct {
	BaseUrl, XGotifyKey string
}

// Send is to send a new request
// & return the response
func (c *Config) Send(r Request) (*http.Response, error) {

	client := &http.Client{}

	request, err := http.NewRequest(c.Method, c.Url, bytes.NewBuffer(c.Body))
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
