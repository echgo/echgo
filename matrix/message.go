package matrix

import (
	"encoding/json"
	"net/url"
)

// CreateMessageBody is to structure the body data
type CreateMessageBody struct {
	Msgtype string `json:"msgtype"`
	Body    string `json:"body"`
}

// CreateMessageReturn is to decode the json data
type CreateMessageReturn struct {
	EventId string `json:"event_id"`
}

// CreateMessage is to create a message on a matrix server
func CreateMessage(body CreateMessageBody, r Request) (CreateMessageReturn, error) {

	address, err := url.JoinPath(r.Domain, "_matrix", "client", "r0", "rooms", r.RoomId, "send", "m.room.message")
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

	parse, err := url.Parse(c.Url)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	newUrl.Add("access_token", r.AccessToken)

	parse.RawQuery = newUrl.Encode()
	c.Url = parse.String()

	response, err := c.Send(r)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	defer response.Body.Close()

	var decode CreateMessageReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CreateMessageReturn{}, err
	}

	return decode, nil

}
