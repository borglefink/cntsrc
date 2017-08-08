// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package find

import (
	"bytes"
	"io/ioutil"
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
		isDir    bool
		filename string
		excluded bool
		expected []byte
	}{
		{true, "dirname1", false, []byte("Directory          dirname1\n")},
		{true, "dirname2", true, []byte("Directory EXCLUDED dirname2\n")},
		{false, "dirname3", false, []byte("File               dirname3\n")},
		{false, "dirname4", true, []byte("File      EXCLUDED dirname4\n")},
	}

	for _, tt := range debugTests {
		var r, w, _ = os.Pipe()

		// redirecting output to a r/w buffer
		testmutex.Lock()
		var oldStdout = os.Stdout
		os.Stdout = w
		showDirectoriesOrFiles(tt.isDir, tt.filename, tt.excluded)

		w.Close()
		os.Stdout = oldStdout
		testmutex.Unlock()

		var actual, _ = ioutil.ReadAll(r)

		if bytes.Compare(actual, tt.expected) != 0 {
			t.Fatalf("Debug info for %s. Expected |%v| but got |%v|", tt.filename, tt.expected, actual)
		}
	}
}
