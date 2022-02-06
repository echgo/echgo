package configuration

// Body is to save & decode the yaml data
type Body struct {
	NotificationChannels NotificationChannels `yaml:"notification_channels"`
	Cronjobs             []Cronjobs           `yaml:"cronjobs"`
}

type NotificationChannels struct {
	Gotify   Gotify   `yaml:"gotify"`
	Telegram Telegram `yaml:"telegram"`
	SMTP     SMTP     `yaml:"smtp"`
	Webhook  Webhook  `yaml:"webhook"`
}

type Gotify struct {
	Domain string `yaml:"domain"`
	Key    string `yaml:"key"`
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

type Cronjobs struct {
	Path         string       `yaml:"path"`
	Cron         string       `yaml:"cron"`
	Notification Notification `yaml:"notification"`
}

type Notification struct {
	Gotify   bool `yaml:"gotify"`
	Telegram bool `yaml:"telegram"`
	SMTP     bool `yaml:"smtp"`
	Webhook  bool `yaml:"webhook"`
}

// Save variables
var (
	filePath = "files/configuration/echgo.yaml"
	Data     Body
)
