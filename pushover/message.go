package pushover

import (
	"encoding/json"
)

// CreateMessageBody is to structure the body data
type CreateMessageBody struct {
	Token   string `json:"token"`
	User    string `json:"user"`
	Message string `json:"message"`
}

// CreateMessageReturn is to decode the json data
type CreateMessageReturn struct {
	Status  int    `json:"status"`
	Request string `json:"request"`
}

// CreateMessage is to create a message to pushover
func CreateMessage(body CreateMessageBody) (CreateMessageReturn, error) {

	convert, err := json.Marshal(body)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	c := Config{
		Path:   "/messages.json",
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
