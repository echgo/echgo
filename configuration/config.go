package configuration

// Body is to save & decode the yaml data
type Body struct {
	Cronjobs []Cronjobs `yaml:"cronjobs"`
}

type Cronjobs struct {
	Path         string       `yaml:"path"`
	Cron         string       `yaml:"cron"`
	Notification Notification `yaml:"notification,omitempty"`
}

type Notification struct {
	Gotify   bool `yaml:"gotify,omitempty"`
	Telegram bool `yaml:"telegram,omitempty"`
	SMTP     bool `yaml:"smtp,omitempty"`
	Webhook  bool `yaml:"webhook,omitempty"`
}

// Save variables
var (
	filePath = "files/configuration/echgo.yaml"
	Data     Body
)
