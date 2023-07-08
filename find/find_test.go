// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package find

import (
	"bytes"
	"io"
	"os"
	"sync"
	"testing"
)

var (
	testmutex = sync.Mutex{}
)

// showDirectoriesOrFiles tests
func TestShowDirectoriesOrFiles(t *testing.T) {
	debug = true
	var debugTests = []struct {
		entryname string
		excluded  bool
		expected  []byte
	}{
		{"entry1", false, []byte("File          entry1\n")},
		{"entry2", true, []byte("File EXCLUDED entry2\n")},
	}

	for _, tt := range debugTests {
		var r, w, _ = os.Pipe()

		// redirecting output to a r/w buffer
		testmutex.Lock()
		var oldStdout = os.Stdout
		os.Stdout = w
		debugIncludedExcludedFiles(tt.entryname, tt.excluded)

		w.Close()
		os.Stdout = oldStdout
		testmutex.Unlock()

		var actual, _ = io.ReadAll(r)

		if !bytes.Equal(actual, tt.expected) {
			t.Fatalf("Debug info for %s. Expected |%v| but got |%v|", tt.entryname, tt.expected, actual)
		}
	}
}
