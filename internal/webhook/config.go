// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package webhook is used to define the configuration
// for the request, the functions so that the
// request can be built and processed.
package webhook

import (
	"bytes"
	"net/http"
)

// channel is to save the channel name for logging.
const channel = "webhook"

// Config is to define config data.
type Config struct {
	Method string
	Body   []byte
}

// Request is to define the request data.
type Request struct {
	Url string
}

// Send is to send a new request & return the response.
func (c *Config) Send(r Request) (*http.Response, error) {

	url := r.Url

	client := &http.Client{}

	request, err := http.NewRequest(c.Method, url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", "echGo")
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil

}
