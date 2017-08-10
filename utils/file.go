// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/kardianos/osext"
)

var (
	windowsNewline       = []byte("\r\n")
	unixNewline          = []byte("\n")
	unixPathSeparator    = "/"
	windowsPathSeparator = "\\"
)

// DetermineNewline determines which kind of newline format the given string contains
func DetermineNewline(stringWithNewline []byte) []byte {
	if bytes.Contains(stringWithNewline, windowsNewline) {
		return windowsNewline
	}

	if bytes.Contains(stringWithNewline, unixNewline) {
		return unixNewline
	}

	return windowsNewline
}

// ResolveStartdir returns the relevant directory to be searched
func ResolveStartdir(pathFromFlag, defaultPath string) string {
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

// IsBinaryFormat determines if the given data represents a binary file
func IsBinaryFormat(data []byte) bool {
	var mimetype = http.DetectContentType(data)
	return strings.Index(mimetype, "text/plain") < 0 &&
		strings.Index(mimetype, "text/html") < 0 &&
		strings.Index(mimetype, "text/xml") < 0
}

// GetExecutableName returns the name of the executable
func GetExecutableName() string {
	filename, err := osext.Executable()

	if err != nil {
		filename, _ = filepath.Abs(os.Args[0])
	}

	return filename
}
