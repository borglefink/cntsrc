// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package utils

import (
	"os"
	"strings"
	"testing"
)

// isInString tests
func TestIsInStringFound(t *testing.T) {
	var actualResult = isInString("Hello", []string{"xx", "He"})
	var expectedResult = true

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestIsInStringNotFound(t *testing.T) {
	var actualResult = isInString("Hello", []string{"xx"})
	var expectedResult = false

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

// isInSlice tests
func TestIsInSliceFound(t *testing.T) {
	var actualResult = isInSlice([]string{"Hello", "Hallo", "Hullu"}, "Hallo")
	var expectedResult = true

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestIsInSliceNotFound(t *testing.T) {
	var actualResult = isInSlice([]string{"Hello", "Hallo", "Hullu"}, "Hilly")
	var expectedResult = false

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestIsWindows(t *testing.T) {
	var actualResult = isWindows()
	var expectedResult = strings.Index(os.Getenv("OS"), "Windows") >= 0

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}
