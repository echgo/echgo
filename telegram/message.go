package telegram

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// CreateMessageReturn is to decode the json data
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

// CreateMessage is to create a message with a bot in a Telegram chat
func CreateMessage(message, chatId, parseMode string, r Request) (CreateMessageReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/sendMessage",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	newUrl.Add("chat_id", chatId)
	newUrl.Add("parse_mode", parseMode)
	newUrl.Add("text", message)

	// Set new url
	parse.RawQuery = newUrl.Encode()
	c.Path = fmt.Sprintf("%s", parse)

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	// Decode response
	var decode CreateMessageReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	// Return data
	return decode, nil

}
