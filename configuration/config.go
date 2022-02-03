package configuration

// ConfigurationBody is to save & decode the yaml data
type ConfigurationBody struct {
	Cronjobs []struct {
		Path         string `yaml:"path"`
		Cron         string `yaml:"cron"`
		Notification struct {
			Gotify   bool `yaml:"gotify,omitempty"`
			Telegram bool `yaml:"telegram,omitempty"`
			SMTP     bool `yaml:"smtp,omitempty"`
			Webhook  bool `yaml:"webhook,omitempty"`
		} `yaml:"notification,omitempty"`
	} `yaml:"cronjobs"`
}

// Save variables
var (
	filePath = "files/configuration/echgo.yaml"
	Data     ConfigurationBody
)
