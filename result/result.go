// Copyright 2014-2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

// Result is metadata for the result
type Result struct {
	Directory          string
	Extensions         map[string]*ExtensionEntry
	TotalNumberOfFiles int
	TotalNumberOfLines int
	TotalSize          int64
}
