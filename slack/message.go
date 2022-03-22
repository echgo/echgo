package slack

import (
	"encoding/json"
	"io"
)

// CreateMessageBody is to structure the body data
type CreateMessageBody struct {
	Text string `json:"text"`
}

// CreateMessage is to create a message on a matrix server
func CreateMessage(body CreateMessageBody, r Request) (string, error) {

	convert, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	c := Config{
		Method: "POST",
		Body:   convert,
	}

	response, err := c.Send(r)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	read, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(read), nil

}
