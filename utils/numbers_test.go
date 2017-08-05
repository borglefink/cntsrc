// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package utils

import "testing"

// int64ToString tests
func TestInt64ToString(t *testing.T) {
	var feed = int64(100000)
	var actualResult = Int64ToString(feed, ' ')
	var expectedResult = "100 000"

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

// round tests
func TestRound(t *testing.T) {
	var roundTests = []struct {
		number    float64
		precision int
		expected  float64
	}{
		{1.45, 1, 1.5},
		{1.44, 1, 1.4},
		{2.45454, 0, 2.0},
		{2.45454, 1, 2.5},
		{2.45454, 2, 2.45},
		{2.45454, 3, 2.455},
		{2.45454, 4, 2.4545},
	}

	for _, tt := range roundTests {
		var actual = Round(tt.number, tt.precision)
		if actual != tt.expected {
			t.Fatalf("Rounding %v with precision %v. Expected %v but got %v", tt.number, tt.precision, tt.expected, actual)
		}
	}
}
