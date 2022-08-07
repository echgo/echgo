package gotify

import (
	"encoding/json"
	"net/url"
	"time"
)

// CreateMessageBody is to structure the body data
type CreateMessageBody struct {
	Priority int                      `json:"priority"`
	Title    string                   `json:"title"`
	Message  string                   `json:"message"`
	Extras   *CreateMessageBodyExtras `json:"extras"`
}

type CreateMessageBodyExtras struct {
	HomeAppliancesLightingOn                  CreateMessageBodyHomeLighting      `json:"home::appliances::lighting::on"`
	HomeAppliancesThermostatChangeTemperature CreateMessageBodyChangeTemperature `json:"home::appliances::thermostat::change_temperature"`
}

type CreateMessageBodyHomeLighting struct {
	Brightness int `json:"brightness"`
}

type CreateMessageBodyChangeTemperature struct {
	Temperature int `json:"temperature"`
}

// CreateMessageReturn is to decode the json data
type CreateMessageReturn struct {
	Id       int    `json:"id"`
	Appid    int    `json:"appid"`
	Message  string `json:"message"`
	Title    string `json:"title"`
	Priority int    `json:"priority"`
	Extras   struct {
		HomeAppliancesLightingOn struct {
			Brightness int `json:"brightness"`
		} `json:"home::appliances::lighting::on"`
		HomeAppliancesThermostatChangeTemperature struct {
			Temperature int `json:"temperature"`
		} `json:"home::appliances::thermostat::change_temperature"`
	} `json:"extras"`
	Date time.Time `json:"date"`
}

// CreateMessage is to create a new message
// To send the request & decode the response
func CreateMessage(body CreateMessageBody, r Request) (CreateMessageReturn, error) {

	address, err := url.JoinPath(r.BaseUrl, "message")
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

	return decode, err

}
