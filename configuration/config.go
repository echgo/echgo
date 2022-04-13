package configuration

// Body is to save & decode the yaml data
type Body struct {
	Channels Channels `yaml:"channels"`
}

type Channels struct {
	Gotify   Gotify   `yaml:"gotify"`
	Matrix   Matrix   `yaml:"matrix"`
	Telegram Telegram `yaml:"telegram"`
	Discord  Discord  `yaml:"discord"`
	Slack    Slack    `yaml:"slack"`
	Trello   Trello   `yaml:"trello"`
	Zendesk  Zendesk  `yaml:"zendesk"`
	OsTicket OsTicket `json:"osticket"`
	Twillo   Twillo   `json:"twillo"`
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

type Discord struct {
	WebhookUrl string `yaml:"webhook_url"`
	BotName    string `yaml:"bot_name"`
}

type Slack struct {
	WebhookUrl string `yaml:"webhook_url"`
}

type Trello struct {
	Key    string `yaml:"key"`
	Token  string `yaml:"token"`
	IdList string `yaml:"id_list"`
}

type Zendesk struct {
	BaseUrl  string `yaml:"base_url"`
	Username string `yaml:"username"`
	ApiToken string `yaml:"api_token"`
}

type OsTicket struct {
	BaseUrl        string `yaml:"base_url"`
	ApiToken       string `yaml:"api_token"`
	Username       string `yaml:"username"`
	EmailRecipient string `yaml:"email_recipient"`
}

type Twillo struct {
	AccountSid        string `yaml:"account_sid"`
	AuthToken         string `yaml:"auth_token"`
	MyPhoneNumber     string `yaml:"my_phone_number"`
	TwilloPhoneNumber string `yaml:"twillo_phone_number"`
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
