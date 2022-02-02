package telegram

import (
	"bytes"
	"net/http"
)

// Set the base url for telegram bot's
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
// & return the response
func (c *Config) Send(r Request) (*http.Response, error) {

	url := baseUrl + r.ApiToken + c.Path

	client := &http.Client{}

	request, err := http.NewRequest(c.Method, url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil

}
