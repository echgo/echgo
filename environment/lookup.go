package environment

import "os"

// Lookup is to check multiple environments
func Lookup(keys ...string) bool {

	for _, value := range keys {
		value = os.Getenv(value)
		if len(value) == 0 {
			return false
		}
	}

	return true

}
