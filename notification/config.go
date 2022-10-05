// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package notification is used to select the different
// file types and to read out and pass on the contents.
package notification

import "path/filepath"

// path is to get the folder path of the notification files.
var (
	path = filepath.Join("files", "notification")
)
