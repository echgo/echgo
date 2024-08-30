// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package discord

import (
	"encoding/json"
)

// CreateMessageBody is to structure the body data.
type CreateMessageBody struct {
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
	Content   string `json:"content"`
}

// CreateMessage is to create a message on a discord server.
func CreateMessage(body CreateMessageBody, r Request) error {

	convert, err := json.Marshal(body)
	if err != nil {
		return err
	}

	c := Config{
		Method: "POST",
		Body:   convert,
	}

	response, err := c.Send(r)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return nil

}
