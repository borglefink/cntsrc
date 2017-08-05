// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

// ExtensionEntry is the file entry for result
type ExtensionEntry struct {
	ExtensionName string
	IsBinary      bool
	NumberOfFiles int32
	NumberOfLines int32
	Filesize      int64
}
