package configuration

import (
	"log"
	"path/filepath"
)

// Body is to save & decode the json data
type Body struct {
	Channels Channels `json:"channels"`
}

type Channels struct {
	Gotify   Gotify   `json:"gotify"`
	Matrix   Matrix   `json:"matrix"`
	Telegram Telegram `json:"telegram"`
	Discord  Discord  `json:"discord"`
	Slack    Slack    `json:"slack"`
	Trello   Trello   `json:"trello"`
	Zendesk  Zendesk  `json:"zendesk"`
	OsTicket OsTicket `json:"osticket"`
	Twillo   Twillo   `json:"twillo"`
	SMTP     SMTP     `json:"smtp"`
	Webhook  Webhook  `json:"webhook"`
}

type Gotify struct {
	Domain string `json:"domain"`
	Key    string `json:"key"`
}

type Matrix struct {
	Domain      string `json:"domain"`
	RoomId      string `json:"room_id"`
	AccessToken string `json:"access_token"`
}

type Telegram struct {
	ApiToken string `json:"api_token"`
	ChatId   string `json:"chat_id"`
}

type Discord struct {
	WebhookUrl string `json:"webhook_url"`
	BotName    string `json:"bot_name"`
}

type Slack struct {
	WebhookUrl string `json:"webhook_url"`
}

type Trello struct {
	Key    string `json:"key"`
	Token  string `json:"token"`
	IdList string `json:"id_list"`
}

type Zendesk struct {
	BaseUrl  string `json:"base_url"`
	Username string `json:"username"`
	ApiToken string `json:"api_token"`
}

type OsTicket struct {
	BaseUrl        string `json:"base_url"`
	ApiToken       string `json:"api_token"`
	Username       string `json:"username"`
	EmailRecipient string `json:"email_recipient"`
}

type Twillo struct {
	AccountSid    string `json:"account_sid"`
	AuthToken     string `json:"auth_token"`
	PhoneNumber   string `json:"phone_number"`
	MyPhoneNumber string `json:"my_phone_number"`
}

type SMTP struct {
	Host           string `json:"host"`
	Port           int    `json:"port"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	EmailRecipient string `json:"email_recipient"`
}

type Webhook struct {
	Domain string `Webhook:"domain"`
}

// path is to save the path of the configuration file
const (
	localPath = "files/configuration/echgo.json"
)

// Data is to save & get the data of the loaded configuration file
var (
	Data Body
)

// absolutePath is to get the absolute file path
func absolutePath() string {

	path, err := filepath.Abs(localPath)
	if err != nil {
		log.Fatalln()
	}

	return path

}
