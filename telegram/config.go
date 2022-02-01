package telegram

import (
	"bytes"
	"net/http"
)

const baseUrl = "https://api.telegram.org/bot"

// Config is to define config data
type Config struct {
	Path, Method string
	ContentType  string
	Body         []byte
	AccessToken  bool
	UploadMedia  bool
}

// Request is to define the request data
type Request struct {
	ApiToken string
}

// Send is to send a new request
func (c *Config) Send(r Request) (*http.Response, error) {

	// Set url
	url := baseUrl + r.ApiToken + c.Path

	// Define client
	client := &http.Client{}

	// Request
	request, err := http.NewRequest(c.Method, url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return data
	return response, nil

}
