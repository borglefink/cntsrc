// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

import "testing"

// InitResult tests
func TestInitResultExtensions(t *testing.T) {
	var extensions = []string{".png", ".exe", ".go"}
	var exclusions = make([]string, 0)
	var actualResult = InitResult(extensions, exclusions, 0)
	var expectedResult = len(extensions)

	if len(actualResult.Extensions) != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestInitResultExceptions(t *testing.T) {
	var extensions = make([]string, 0)
	var exclusions = []string{"//dir", "//bin//", ".gitignore"}
	var actualResult = InitResult(extensions, exclusions, 0)
	var expectedResult = len(exclusions)

	if len(actualResult.Exclusions) != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestInitResultBigFiles(t *testing.T) {
	var extensions = make([]string, 0)
	var exclusions = make([]string, 0)
	var actualResult = InitResult(extensions, exclusions, 5)
	var expectedResult = 5

	if actualResult.NumberOfBigFiles != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}
