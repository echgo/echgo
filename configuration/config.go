package configuration

// Body is to save & decode the yaml data
type Body struct {
	Channels Channels `yaml:"channels"`
}

type Channels struct {
	Gotify   Gotify   `yaml:"gotify"`
	Matrix   Matrix   `yaml:"matrix"`
	Telegram Telegram `yaml:"telegram"`
	SMTP     SMTP     `yaml:"smtp"`
	Webhook  Webhook  `yaml:"webhook"`
}

type Gotify struct {
	Domain string `yaml:"domain"`
	Key    string `yaml:"key"`
}

type Matrix struct {
	Domain      string `yaml:"domain"`
	RoomId      string `yaml:"room_id"`
	AccessToken string `yaml:"access_token"`
}

type Telegram struct {
	ApiToken string `yaml:"api_token"`
	ChatId   string `yaml:"chat_id"`
}

type SMTP struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	EmailRecipient string `yaml:"email_recipient"`
}

type Webhook struct {
	Domain string `Webhook:"domain"`
}

// path is to save the path of the configuration file
const (
	path = "files/configuration/echgo.yaml"
)

// Data is to save & get the data of the loaded configuration file
var (
	Data Body
)
