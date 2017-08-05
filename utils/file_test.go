// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package utils

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

const (
	notexistdirectory = "thisshouldnotbeasubdirectoryinthecurrentdirectory"
)

// determineNewline tests
func TestDetermineWindowsNewline(t *testing.T) {
	var actualResult = DetermineNewline([]byte("Hello\r\n"))
	var expectedResult = windowsNewline

	if bytes.Compare(actualResult, expectedResult) != 0 {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestDetermineUnixNewline(t *testing.T) {
	var actualResult = DetermineNewline([]byte("Hello\n"))
	var expectedResult = unixNewline

	if bytes.Compare(actualResult, expectedResult) != 0 {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestDetermineDefaultWindowsNewline(t *testing.T) {
	var actualResult = DetermineNewline([]byte("Hello"))
	var expectedResult = windowsNewline

	if bytes.Compare(actualResult, expectedResult) != 0 {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

// isBinaryFormat tests
func TestIsBinaryFormatFalse(t *testing.T) {
	var actualResult = IsBinaryFormat([]byte("<html><body><br/></body></html>"))
	var expectedResult = false

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestIsBinaryFormatTrue(t *testing.T) {
	var actualResult = IsBinaryFormat([]byte("‰PNG IHDR  h     ‰"))
	var expectedResult = true

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

//func getDirectory(pathFromFlag, defaultPath string) string {
func TestGetDirectoryNotExist(t *testing.T) {
	if os.Getenv("BE_GETDIRECTORY") == "1" {
		getDirectory(notexistdirectory, "")
		return
	}

	var cmd = exec.Command(os.Args[0], "-test.run=TestGetDirectoryNotExist")
	cmd.Env = append(os.Environ(), "BE_GETDIRECTORY=1")

	var err = cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	t.Fatalf("Expected error but got %v", err)
}
