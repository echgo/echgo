package webhook

import "encoding/json"

// SendBody is to structure the body data
type SendBody struct {
	Message string `json:"message"`
}

// Send is to send the webhook
func Send(body SendBody, r Request) error {

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
