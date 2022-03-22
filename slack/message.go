package slack

import "encoding/json"

// CreateMessageBody is to structure the body data
type CreateMessageBody struct {
	Text string `json:"text"`
}

// CreateMessage is to create a message on a matrix server
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
