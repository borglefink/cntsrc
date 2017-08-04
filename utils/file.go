// Copyright 2014-2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	windowsNewline       = "\r\n"
	unixNewline          = "\n"
	oldMacNewline        = "\r"
	unixPathSeparator    = "/"
	windowsPathSeparator = "\\"
)

// determineNewline
func determineNewline(stringWithNewline string) string {
	if strings.Contains(stringWithNewline, windowsNewline) {
		return windowsNewline
	}

	if strings.Contains(stringWithNewline, unixNewline) {
		return unixNewline
	}

	if strings.Contains(stringWithNewline, oldMacNewline) {
		return oldMacNewline
	}

	return windowsNewline
}

// getDirectory
func getDirectory(pathFromFlag, defaultPath string) string {
	var err error

	// First non-flag argument should be the starting directory
	var path = pathFromFlag

	// If no directory given, use the current directory
	if len(path) == 0 {
		path = filepath.Dir(defaultPath)
	}

	// Getting the full path, if necessary
	path, err = filepath.Abs(path)
	if err != nil {
		fmt.Printf("Directory [%v] does not exist.\n", path)
		os.Exit(1)
	}

	// Removing quotes, if any
	path = strings.Replace(path, "\"", "", -1)

	// Checking if directory is ok
	_, err = os.Stat(path)
	if err != nil {
		fmt.Printf("Directory [%v] does not exist.\n", path)
		os.Exit(1)
	}

	return path
}

// getPathSeparator
func getPathSeparator() string {
	if isWindows() {
		return windowsPathSeparator
	}
	return unixPathSeparator
}

// isBinaryFormat
func isBinaryFormat(data []byte) bool {
	var mimetype = http.DetectContentType(data)
	return strings.Index(mimetype, "text/plain") < 0 &&
		strings.Index(mimetype, "text/html") < 0 &&
		strings.Index(mimetype, "text/xml") < 0
}
