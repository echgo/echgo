// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package environment

import (
	"os"
	"strconv"
)

// Integer is to get look up the environment and return is as integer.
func Integer(key string) int {

	value, ok := os.LookupEnv(key)
	if ok {
		integer, err := strconv.Atoi(value)
		if err != nil {
			return 0
		}
		return integer
	}

	return 0

}
