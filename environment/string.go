package environment

import (
	"os"
)

// String is to get look up the environment and return is as string
func String(key string) string {

	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	return ""

}
