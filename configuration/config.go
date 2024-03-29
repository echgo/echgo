// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package configuration is used to import or create the
// different configurations for the different channels,
// so that a configuration json file is used
// to store the different access data.
package configuration

// Body is to save & decode the json data.
type Body struct {
	Channels Channels `json:"channels"`
}

type Channels struct {
	Gotify   Gotify   `json:"gotify"`
	Pushover Pushover `json:"pushover"`
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
	BaseUrl string `json:"base_url"`
	Key     string `json:"key"`
}

type Pushover struct {
	Token string `json:"token"`
	User  string `json:"user"`
}

type Matrix struct {
	BaseUrl     string `json:"base_url"`
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
	Url string `json:"url"`
}

// Data is to save & get the data of the loaded configuration file.
var (
	Data Body
)
