package telegram

import (
	"bytes"
	"net/http"
)

// Set the base url for telegram bot's
const baseUrl = "https://api.telegram.org"

// Config is to define config data
type Config struct {
	Url, Method string
	Body        []byte
}

// Request is to define the request data
type Request struct {
	ApiToken string
	ChatId   string
}

// Send is to send a new request
// & return the response
func (c *Config) Send(r Request) (*http.Response, error) {

	client := &http.Client{}

	request, err := http.NewRequest(c.Method, c.Url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil

}
