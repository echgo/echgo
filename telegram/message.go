// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package telegram

import (
	"encoding/json"
	"net/url"
)

// CreateMessageReturn is to decode the json data.
type CreateMessageReturn struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			Id                          int    `json:"id"`
			Title                       string `json:"title"`
			Type                        string `json:"type"`
			AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}

// CreateMessage is to create a message with a bot in a telegram chat.
func CreateMessage(message, parseMode string, r Request) (CreateMessageReturn, error) {

	address, err := url.JoinPath(baseUrl, r.ApiToken, "sendMessage")
	if err != nil {
		return CreateMessageReturn{}, err
	}

	c := Config{
		Url:    address,
		Method: "GET",
		Body:   nil,
	}

	parse, err := url.Parse(c.Url)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	newUrl.Add("chat_id", r.ChatId)
	newUrl.Add("parse_mode", parseMode)
	newUrl.Add("text", message)

	parse.RawQuery = newUrl.Encode()
	c.Url = parse.String()

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
