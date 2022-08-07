package trello

import (
	"net/http"
)

// Set the base url for trello
const baseUrl = "https://api.trello.com"

// Config is to define config data
type Config struct {
	Url, Method string
}

// Request is to define the request data
type Request struct {
	Key    string
	Token  string
	IdList string
}

// Send is to send a new request
// & return the response
func (c *Config) Send() (*http.Response, error) {

	client := &http.Client{}

	request, err := http.NewRequest(c.Method, c.Url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil

}
