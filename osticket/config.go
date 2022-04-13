package osticket

import (
	"bytes"
	"net/http"
)

// Config is to define config data
type Config struct {
	Path, Method string
	Body         []byte
}

// Request is to define the request data
type Request struct {
	BaseUrl  string
	ApiToken string
}

// Send is to send a new request
// & return the response
func (c *Config) Send(r Request) (*http.Response, error) {

	url := r.BaseUrl + c.Path

	client := &http.Client{}

	request, err := http.NewRequest(c.Method, url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", r.ApiToken)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil

}
