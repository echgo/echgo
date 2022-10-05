// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package console is used to output the various
// functions in the command line interface.
package console

import (
	"fmt"
	"strings"
)

// Initialize is to create an initialized text.
func Initialize() {

	fmt.Printf("\n%s\n", strings.Repeat("*", 80))
	fmt.Printf("Welcome to echGo! Here you will find the most important information\ndirectly in the console. For more information visit https://echgo.org.\n")
	fmt.Printf("%s\n\n", strings.Repeat("*", 80))

}
