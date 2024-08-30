// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pushover

import (
	"encoding/json"
	"net/url"
)

// CreateMessageBody is to structure the body data.
type CreateMessageBody struct {
	Token   string `json:"token"`
	User    string `json:"user"`
	Message string `json:"message"`
}

// CreateMessageReturn is to decode the json data.
type CreateMessageReturn struct {
	Status  int    `json:"status"`
	Request string `json:"request"`
}

// CreateMessage is to create a message to pushover.
func CreateMessage(body CreateMessageBody) (CreateMessageReturn, error) {

	address, err := url.JoinPath(baseUrl, "1", "messages.json")
	if err != nil {
		return CreateMessageReturn{}, err
	}

	convert, err := json.Marshal(body)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	c := Config{
		Url:    address,
		Method: "POST",
		Body:   convert,
	}

	response, err := c.Send()
	if err != nil {
		return CreateMessageReturn{}, err
	}

	var decode CreateMessageReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	return decode, nil

}
