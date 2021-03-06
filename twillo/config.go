package twillo

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// Set the base url for twillo
const baseUrl = "https://api.twilio.com/2010-04-01"

// Config is to define config data
type Config struct {
	Path, Method string
	Body         string
}

// Request is to define the request data
type Request struct {
	AccountSid string
	AuthToken  string
}

// Send is to send a new request
// & return the response
func (c *Config) Send(r Request) (*http.Response, error) {

	url := baseUrl + c.Path

	encoded := base64.StdEncoding.EncodeToString([]byte(r.AccountSid + ":" + r.AuthToken))

	client := &http.Client{}

	request, err := http.NewRequest(c.Method, url, strings.NewReader(c.Body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", "Basic "+encoded)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil

}
