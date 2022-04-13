package twillo

import (
	"encoding/json"
	"net/url"
)

// CreateMessageBody is to structure the body data
type CreateMessageBody struct {
	Message           string
	MyPhoneNumber     string
	TwilloPhoneNumber string
}

// CreateMessageReturn is to decode the json data
type CreateMessageReturn struct {
	Sid                 string      `json:"sid"`
	DateCreated         string      `json:"date_created"`
	DateUpdated         string      `json:"date_updated"`
	DateSent            interface{} `json:"date_sent"`
	AccountSid          string      `json:"account_sid"`
	To                  string      `json:"to"`
	From                string      `json:"from"`
	MessagingServiceSid interface{} `json:"messaging_service_sid"`
	Body                string      `json:"body"`
	Status              interface{} `json:"status"`
	NumSegments         string      `json:"num_segments"`
	NumMedia            string      `json:"num_media"`
	Direction           string      `json:"direction"`
	ApiVersion          string      `json:"api_version"`
	Price               interface{} `json:"price"`
	PriceUnit           string      `json:"price_unit"`
	ErrorCode           interface{} `json:"error_code"`
	ErrorMessage        interface{} `json:"error_message"`
	Uri                 string      `json:"uri"`
	SubresourceUris     struct {
		Media string `json:"media"`
	} `json:"subresource_uris"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
}

// CreateMessage is to create a message via twillo
func CreateMessage(body CreateMessageBody, r Request) (CreateMessageReturn, error) {

	data := url.Values{}

	data.Set("Body", body.Message)
	data.Set("To", body.MyPhoneNumber)
	data.Set("From", body.TwilloPhoneNumber)

	encode := data.Encode()

	c := Config{
		Path:   "/Accounts/" + r.AccountSid + "/Messages.json",
		Method: "POST",
		Body:   encode,
	}

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
