// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package find

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/MichaelTJones/walk"

	"cntsrc/config"
	"cntsrc/print"
	"cntsrc/result"
	"cntsrc/utils"
)

const (
	// bigFileMaxLength is the maximum length for name of a "big file"
	bigFileMaxLength = 56
)

var (
	res                  result.Result
	startdir             = ""
	debug                = false
	currentPathSeparator = string(os.PathSeparator)
	mutex                = sync.Mutex{}
)

// SetDebug turns on (or off) showing of included/excluded
// files/directories as a debugging facility for the
// exclusions in the config file
func SetDebug(showDebug bool) {
	debug = showDebug

	if debug {
		print.DebugHeader()
	}
}

// isExcluded
func isExcluded(filename string) bool {
	var fulldir, _ = filepath.Abs(filename)

	return utils.IsInString(fulldir, res.Exclusions)
}

// debugIncludedExcludedFiles
func debugIncludedExcludedFiles(filename string, excluded bool) {
	var status = "        "
	if excluded {
		status = "EXCLUDED"
	}

	fmt.Printf("File %s %s\n", status, strings.Replace(filename, startdir+currentPathSeparator, "", 1))
}

// forEachFileSystemEntry
func forEachFileSystemEntry(filename string, f os.FileInfo, err error) error {
	if f == nil {
		return nil
	}

	// Ignore directories
	if f.IsDir() {
		return nil
	}

	var excluded = isExcluded(filename)

	if debug {
		debugIncludedExcludedFiles(filename, excluded)
	}

	if excluded {
		return nil
	}

	var ext = filepath.Ext(filename)

	var entry, willBeCounted = res.Extensions[ext]
	if !willBeCounted {
		return nil
	}

	atomic.AddInt32(&(entry.NumberOfFiles), 1)
	atomic.AddInt32(&res.TotalNumberOfFiles, 1)

	var size = f.Size()
	atomic.AddInt64(&(entry.Filesize), size)
	atomic.AddInt64(&res.TotalSize, size)

	// Slurp the whole file into memory
	var contents, readerr = os.ReadFile(filename)

	if readerr != nil {
		if debug {
			fmt.Printf("Problem reading inputfile %s, error:%v\n", filename, readerr)
		}
		return readerr
	}

	var isBinary = utils.IsBinaryFormat(contents)

	// Binary files will not have "number of lines", but
	// will need to have the binary flag set for the report
	if isBinary {
		mutex.Lock()
		if !entry.IsBinary {
			entry.IsBinary = true
		}
		mutex.Unlock()
		return nil
	}

	var newline = utils.DetermineNewline(contents)
	var numberOfLines = int32(len(bytes.Split(contents, []byte(newline))))

	atomic.AddInt32(&(entry.NumberOfLines), numberOfLines)
	atomic.AddInt32(&res.TotalNumberOfLines, numberOfLines)

	if res.NumberOfBigFiles > 0 {
		mutex.Lock()
		res.BigFiles = append(res.BigFiles, result.FileSize{
			Name:  truncBigFileName(startdir, filename),
			Size:  size,
			Lines: numberOfLines,
		})
		mutex.Unlock()
	}

	return nil
}

// truncBigFileName
func truncBigFileName(startdir, bigFileName string) string {
	// remove startdir path
	var startdirpos = strings.Index(bigFileName, startdir)
	if startdirpos >= 0 {
		bigFileName = bigFileName[len(startdir)+1:]
	}

	// possibly truncate name and add "..." in front
	var bigFileNameLength = len(bigFileName)
	if bigFileNameLength > bigFileMaxLength {
		bigFileName = bigFileName[(bigFileNameLength - bigFileMaxLength):]
		bigFileName = "..." + bigFileName[3:]
	}

	return bigFileName
}

// The function All searches the given directory and returns the search result
func All(dir string, cfg config.Config, bigFiles int) result.Result {
	startdir = dir
	res = result.InitResult(cfg.FileExtensions, cfg.Exclusions, bigFiles)

	var err = walk.Walk(startdir, forEachFileSystemEntry)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}

	return res
}
