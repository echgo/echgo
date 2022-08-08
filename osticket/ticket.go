package osticket

import (
	"encoding/json"
	"io"
	"net/url"
)

// CreateTicketBody is to structure the body data
type CreateTicketBody struct {
	Alert       bool                `json:"alert"`
	Autorespond bool                `json:"autorespond"`
	Source      string              `json:"source"`
	Name        string              `json:"name"`
	Email       string              `json:"email"`
	Phone       string              `json:"phone"`
	Subject     string              `json:"subject"`
	Ip          string              `json:"ip"`
	Message     string              `json:"message"`
	Attachments []map[string]string `json:"attachments,omitempty"`
}

// CreateTicket is to create a ticket in zendesk
func CreateTicket(body CreateTicketBody, r Request) (string, error) {

	address, err := url.JoinPath(r.BaseUrl, "api", "tickets.json")
	if err != nil {
		return "", err
	}

	convert, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	c := Config{
		Url:    address,
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
