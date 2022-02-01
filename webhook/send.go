package webhook

import "encoding/json"

// SendBody is to structure the body data
type SendBody struct {
	Headline string `json:"headline"`
	Message  string `json:"message"`
}

// Send is to send the webhook
func Send(body SendBody, r Request) error {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Set config for request
	c := Config{
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return err
	}

	// Close request
	defer response.Body.Close()

	// Return nothing
	return nil

}
