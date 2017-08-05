// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package utils

import (
	"os"
	"strings"
)

// IsInString returns true if the string is found in the target
func IsInString(stringToSearch string, stringsToSearchFor []string) bool {
	var isFound = false

	for _, searchItem := range stringsToSearchFor {
		if strings.Contains(stringToSearch, searchItem) {
			isFound = true
			break
		}
	}

	return isFound
}

// isInSlice
func isInSlice(sliceToSearch []string, stringToSearchFor string) bool {
	var isFound = false

	for _, searchItem := range sliceToSearch {
		if searchItem == stringToSearchFor {
			isFound = true
			break
		}
	}

	return isFound
}

// isWindows
func isWindows() bool {
	return strings.Index(os.Getenv("OS"), "Windows") >= 0
}
