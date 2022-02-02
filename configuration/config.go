package configuration

// ConfigurationBody is to save & decode the yaml data
type ConfigurationBody struct {
	Cronjobs []struct {
		Name string `yaml:"name"`
		Cron string `yaml:"cron"`
	} `yaml:"cronjobs"`
}

// Save variables
var (
	filePath = "files/configuration/echgo.yaml"
	Data     ConfigurationBody
)
