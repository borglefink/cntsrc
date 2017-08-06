// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package print

import (
	"bytes"
	"io/ioutil"
	"os"
	"sync"
	"testing"

	"cntsrc/result"
)

var (
	mutex = sync.Mutex{}
)

// printEntry tests
func TestPrintEntrySource(t *testing.T) {
	var extensionEntry = &result.ExtensionEntry{
		ExtensionName: ".go",
		Filesize:      1000,
		NumberOfFiles: 1,
		NumberOfLines: 1000,
	}
	var r, w, _ = os.Pipe()

	// redirecting output to a r/w buffer
	mutex.Lock()
	var oldStdout = os.Stdout
	os.Stdout = w
	printEntry(extensionEntry, int32(20), int64(1000))
	w.Close()
	os.Stdout = oldStdout
	mutex.Unlock()

	var actualResult, _ = ioutil.ReadAll(r)
	var expectedResult = []byte(".go                  1        1 000 5000.0         1 000  100.0\n")

	if bytes.Compare(actualResult, expectedResult) != 0 {
		t.Fatalf("Expected |%s| but got |%s|", expectedResult, actualResult)
	}
}

func TestPrintEntryBinary(t *testing.T) {
	var extensionEntry = &result.ExtensionEntry{
		ExtensionName: ".png",
		IsBinary:      true,
		Filesize:      1000,
		NumberOfFiles: 1,
	}
	var r, w, _ = os.Pipe()

	// redirecting output to a r/w buffer
	mutex.Lock()
	var oldStdout = os.Stdout
	os.Stdout = w
	printEntry(extensionEntry, 0, int64(1000))
	w.Close()
	os.Stdout = oldStdout
	mutex.Unlock()

	var actualResult, _ = ioutil.ReadAll(r)
	var expectedResult = []byte(".png                 1                             1 000  100.0\n")

	if bytes.Compare(actualResult, expectedResult) != 0 {
		t.Fatalf("Expected |%s| but got |%s|", expectedResult, actualResult)
	}
}

// printBigFiles tests
func TestPrintBigFiles(t *testing.T) {
	var res = result.Result{
		NumberOfBigFiles: 3,
		BigFiles: []result.FileSize{
			{Name: "File1", Size: 1000, Lines: 10},
			{Name: "File2", Size: 3000, Lines: 30},
			{Name: "File3", Size: 2000, Lines: 20},
			{Name: "File4", Size: 500, Lines: 5},
		},
	}
	var r, w, _ = os.Pipe()

	// redirecting output to a r/w buffer
	mutex.Lock()
	var oldStdout = os.Stdout
	os.Stdout = w
	printBigFiles(res)
	w.Close()
	os.Stdout = oldStdout
	mutex.Unlock()

	var actualResult, _ = ioutil.ReadAll(r)
	var expectedResult = []byte(`

The   3 largest files are:                 #lines
-------------------------------------------------
File2                                          30
File3                                          20
File1                                          10
`)

	if bytes.Compare(actualResult, expectedResult) != 0 {
		t.Fatalf("Expected |%s| but got |%s|", expectedResult, actualResult)
	}
}

// printFooter tests
func TestPrintFooter(t *testing.T) {
	var res = result.Result{
		TotalSize:          1,
		TotalNumberOfFiles: 1,
		TotalNumberOfLines: 1,
	}
	var r, w, _ = os.Pipe()

	// redirecting output to a r/w buffer
	mutex.Lock()
	var oldStdout = os.Stdout
	os.Stdout = w
	printFooter(res)
	w.Close()
	os.Stdout = oldStdout
	mutex.Unlock()

	var actualResult, _ = ioutil.ReadAll(r)
	var expectedResult = []byte(
		`---------------------------------------------------------------
Total:               1            1  100.0             1  100.0
`)

	if bytes.Compare(actualResult, expectedResult) != 0 {
		t.Fatalf("Expected |%s| but got |%s|", expectedResult, actualResult)
	}
}

func TestPrintFooterNoFiles(t *testing.T) {
	var res = result.Result{}
	var r, w, _ = os.Pipe()

	// redirecting output to a r/w buffer
	mutex.Lock()
	var oldStdout = os.Stdout
	os.Stdout = w
	printFooter(res)
	w.Close()
	os.Stdout = oldStdout
	mutex.Unlock()

	var actualResult, _ = ioutil.ReadAll(r)
	var expectedResult = []byte(
		`No files found.

Check given directory,or maybe 
check extensions in config file.
---------------------------------------------------------------
`)

	if bytes.Compare(actualResult, expectedResult) != 0 {
		t.Fatalf("Expected |%s| but got |%s|", expectedResult, actualResult)
	}
}

/*



// Result displays the result from the search
func Result(startdir string, res result.Result) {
	// Show result header
	printHeader(startdir)

	// Sorting keys for presentation
	var keys, binaryKeys = getKeys(res.Extensions)

	// Show result for sourcecode, only for extensions found
	for _, ext := range keys {
		printEntry(res.Extensions[ext], res.TotalNumberOfLines, res.TotalSize)
	}

	// For convenience, show result for binaries separately
	for _, ext := range binaryKeys {
		printEntry(res.Extensions[ext], 0, res.TotalSize)
	}

	// Show footer
	printFooter(res)

	if len(res.BigFiles) > 0 {
		printBigFiles(res)
	}
}
*/
