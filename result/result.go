// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

// Result contains all data and totals for the count
type Result struct {
	Directory          string
	Extensions         map[string]*ExtensionEntry
	TotalNumberOfFiles int32
	TotalNumberOfLines int32
	TotalSize          int64
	Exclusions         []string
	NumberOfBigFiles   int
	BigFiles           FileSizes
}

// InitResult initialises the result
func InitResult(extensions []string, exclusions []string, bigFiles int) Result {
	var r = Result{
		Extensions:       make(map[string]*ExtensionEntry),
		Exclusions:       exclusions,
		NumberOfBigFiles: bigFiles,
	}

	for _, ext := range extensions {
		r.Extensions[ext] = &ExtensionEntry{
			ExtensionName: ext,
		}
	}

	return r
}
