// Copyright 2014-2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package result

// FileSize contains file size
type FileSize struct {
	Name  string
	Size  int64
	Lines int
}

// FileSizes contains a slice of FileSize
type FileSizes []FileSize

//func (p FileSizes) Add(name string, size int64) { p = append(p, FileSize{name, size}) }
func (p FileSizes) Len() int           { return len(p) }
func (p FileSizes) Less(i, j int) bool { return p[i].Lines > p[j].Lines }
func (p FileSizes) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
