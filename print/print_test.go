// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package print

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"sync"
	"testing"

	"cntsrc/result"
)

var (
	mutex = sync.Mutex{}
)

// printHeader tests
func TestPrintHeader(t *testing.T) {
	var r, w, _ = os.Pipe()

	// redirecting output to a r/w buffer
	mutex.Lock()
	var oldStdout = os.Stdout
	os.Stdout = w
	printHeader("/home/user/code/project")
	w.Close()
	os.Stdout = oldStdout
	mutex.Unlock()

	var actualResult, _ = ioutil.ReadAll(r)
	var expectedResult = []byte(`
Directory processed:
/home/user/code/project
---------------------------------------------------------------
filetype        #files       #lines  line%          size  size%
---------------------------------------------------------------
`)

	if bytes.Compare(actualResult, expectedResult) != 0 {
		t.Fatalf("Expected |%s| but got |%s|", expectedResult, actualResult)
	}
}

// printEntry tests
func TestPrintEntrySource(t *testing.T) {
	var extensionEntry = {
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
	var extensionEntry = {
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

The 3 largest files are:                                 #lines
---------------------------------------------------------------

File2                                                        30
File3                                                        20
File1                                                        10
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

// getKeys tests
func TestGetKeys(t *testing.T) {
	//var extensions = make(map[string]*result.ExtensionEntry) //[]string{".png", ".exe", ".go"}
	var extensions = map[string]*result.ExtensionEntry{
		".xml":  {NumberOfFiles: 1, IsBinary: false},
		".js":   {NumberOfFiles: 1, IsBinary: false},
		".go":   {NumberOfFiles: 1, IsBinary: false},
		".css":  {NumberOfFiles: 0, IsBinary: false},
		".html": {NumberOfFiles: 0, IsBinary: false},
		".zip":  {NumberOfFiles: 1, IsBinary: true},
		".png":  {NumberOfFiles: 1, IsBinary: true},
		".exe":  {NumberOfFiles: 0, IsBinary: true},
		"":      {NumberOfFiles: 1, IsBinary: true},
	}
	var actualResult1, actualResult2 = getKeys(extensions)
	var expectedResult1 = []string{".go", ".js", ".xml"}
	var expectedResult2 = []string{"", ".png", ".zip"}

	if !reflect.DeepEqual(actualResult1, expectedResult1) {
		t.Fatalf("1.Expected %v but got %v", expectedResult1, actualResult1)
	}

	if !reflect.DeepEqual(actualResult2, expectedResult2) {
		t.Fatalf("2.Expected %v but got %v", expectedResult2, actualResult2)
	}
}

// Result tests
func TestResult(t *testing.T) {
	var res = result.Result{
		Extensions: map[string]*result.ExtensionEntry{
			".xml":  {ExtensionName: ".xml", NumberOfFiles: 1, NumberOfLines: 10, Filesize: 200, IsBinary: false},
			".js":   {ExtensionName: ".js", NumberOfFiles: 1, NumberOfLines: 10, Filesize: 200, IsBinary: false},
			".go":   {ExtensionName: ".go", NumberOfFiles: 10, NumberOfLines: 100, Filesize: 2000, IsBinary: false},
			".css":  {ExtensionName: ".css", NumberOfFiles: 0, NumberOfLines: 0, Filesize: 0, IsBinary: false},
			".html": {ExtensionName: ".html", NumberOfFiles: 0, NumberOfLines: 0, Filesize: 0, IsBinary: false},
			".zip":  {ExtensionName: ".zip", NumberOfFiles: 1, Filesize: 200, IsBinary: true},
			".png":  {ExtensionName: ".png", NumberOfFiles: 1, Filesize: 200, IsBinary: true},
			".exe":  {ExtensionName: ".exe", NumberOfFiles: 0, Filesize: 0, IsBinary: true},
		},
		TotalNumberOfFiles: 14,
		TotalNumberOfLines: 120,
		TotalSize:          2800,
		NumberOfBigFiles:   3,
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
	Result("/home/user/code/project", res)
	w.Close()
	os.Stdout = oldStdout
	mutex.Unlock()

	var actualResult, _ = ioutil.ReadAll(r)
	var expectedResult = []byte(`
Directory processed:
/home/user/code/project
-------------------------------------------------------------------
filetype            #files       #lines  line%          size  size%
-------------------------------------------------------------------
.go                     10          100   83.3         2 000   71.4
.js                      1           10    8.3           200    7.1
.xml                     1           10    8.3           200    7.1
.png                     1                               200    7.1
.zip                     1                               200    7.1
-------------------------------------------------------------------
Total:                  14          120  100.0         2 800  100.0


The 3 largest files are:                                 #lines
---------------------------------------------------------------

File2                                                        30
File3                                                        20
File1                                                        10
`)

	if bytes.Compare(actualResult, expectedResult) != 0 {
		t.Fatalf("Expected |%s| but got |%s|", expectedResult, actualResult)
	}
}
