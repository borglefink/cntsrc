// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package utils

import (
	"bytes"
	"math"
	"strconv"
)

// Int64ToString shows a string representation of a big int
func Int64ToString(n int64, separator rune) string {
	var s = strconv.FormatInt(n, 10)
	var startOffset = 0
	var buff bytes.Buffer

	if n < 0 {
		startOffset = 1
		buff.WriteByte('-')
	}

	var length = len(s)
	var commaIndex = 3 - ((length - startOffset) % 3)

	if commaIndex == 3 {
		commaIndex = 0
	}

	for i := startOffset; i < length; i++ {
		if commaIndex == 3 {
			buff.WriteRune(separator)
			commaIndex = 0
		}
		commaIndex++
		buff.WriteByte(s[i])
	}

	return buff.String()
}

// Round return rounded version of x with prec precision.
// http://grokbase.com/t/gg/golang-nuts/12ag1s0t5y/go-nuts-re-function-round
func Round(x float64, prec int) float64 {
	var rounder float64
	var pow = math.Pow(10, float64(prec))
	var intermed = x * pow
	var _, frac = math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
