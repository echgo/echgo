package zendesk

import (
	"bytes"
	"encoding/base64"
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
	Username string
	ApiToken string
}

// Send is to send a new request
// & return the response
func (c *Config) Send(r Request) (*http.Response, error) {

	url := r.BaseUrl + c.Path

	encoded := base64.StdEncoding.EncodeToString([]byte(r.Username + "/token:" + r.ApiToken))

	client := &http.Client{}

	request, err := http.NewRequest(c.Method, url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic "+encoded)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil

}
