// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package utils

import "strings"

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
