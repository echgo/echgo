// Copyright 2024 Jonas Kwiedor. All rights reserved.

// Package environment is used to check the environment
// variables for availability and to return them in different
// data types, so that they can be used directly converted.
package environment

import (
	"os"
	"strconv"
)

// Boolean is to get look up the environment and return is as boolean.
func Boolean(key string) bool {

	value, ok := os.LookupEnv(key)
	if ok {
		boolean, err := strconv.ParseBool(value)
		if err != nil {
			return false
		}
		return boolean
	}

	return false

}
