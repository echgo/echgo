// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package twillo is used to define the configuration
// for the request, the functions so that the
// request can be built and processed.
package twillo

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// baseUrl is to set the base url for the request.
// channel is to save the channel name for logging.
const (
	baseUrl = "https://api.twilio.com"
	channel = "twillo"
)

// Config is to define config data.
type Config struct {
	Url, Method string
	Body        string
}

// Request is to define the request data.
type Request struct {
	AccountSid string
	AuthToken  string
}

// Send is to send a new request & return the response.
func (c *Config) Send(r Request) (*http.Response, error) {

	encoded := base64.StdEncoding.EncodeToString([]byte(r.AccountSid + ":" + r.AuthToken))

	client := &http.Client{}

	request, err := http.NewRequest(c.Method, c.Url, strings.NewReader(c.Body))
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
