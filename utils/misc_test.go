// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package utils

import "testing"

// IsInString tests
func TestIsInStringFound(t *testing.T) {
	var actualResult = IsInString("Hello", []string{"xx", "He"})
	var expectedResult = true

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestIsInStringNotFound(t *testing.T) {
	var actualResult = IsInString("Hello", []string{"xx"})
	var expectedResult = false

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}
