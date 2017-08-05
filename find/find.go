// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package find

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/MichaelTJones/walk"

	"cntsrc/config"
	"cntsrc/result"
	"cntsrc/utils"
)

var (
	res                  result.Result
	startdir             = ""
	debug                = false
	numberOfBigFiles     = 0
	currentPathSeparator = string(os.PathSeparator)
	mutex                = sync.Mutex{}
)

// SetDebug turns on (or off) showing of
// included/excluded files/directories
// as a debugging facility for the config file
func SetDebug(showDebug bool) {
	debug = showDebug
}

// SetBigFiles sets the number of big files to be shown
func SetBigFiles(showBigFiles int) {
	numberOfBigFiles = showBigFiles
}

// isExcluded
func isExcluded(filename string) bool {
	var fulldir, _ = filepath.Abs(filename)

	return utils.IsInString(fulldir, res.Exclusions)
}

// showDirectoriesOrFiles
func showDirectoriesOrFiles(isDir bool, filename string, excluded bool) {
	if !debug {
		return
	}

	var prompt string
	if isDir {
		prompt = "Directory"
	} else {
		prompt = "File     "
	}

	var status string
	if excluded {
		status = " EXCLUDED "
	} else {
		status = "          "
	}

	fmt.Printf("%s %s%s\n", prompt, status, strings.Replace(filename, startdir+currentPathSeparator, "", 1))
}

// forEachFileSystemEntry
func forEachFileSystemEntry(filename string, f os.FileInfo, err error) error {
	if f == nil {
		return nil
	}

	var excluded = isExcluded(filename)

	if debug {
		showDirectoriesOrFiles(f.IsDir(), filename, excluded)
	}

	if f.IsDir() {
		return nil
	}

	if !excluded {
		var ext = filepath.Ext(filename)

		var _, willBeCounted = res.Extensions[ext]
		if !willBeCounted {
			return nil
		}

		atomic.AddInt32(&(res.Extensions[ext].NumberOfFiles), 1)
		atomic.AddInt32(&res.TotalNumberOfFiles, 1)

		var size = f.Size()
		atomic.AddInt64(&(res.Extensions[ext].Filesize), size)
		atomic.AddInt64(&res.TotalSize, size)

		// Slurp the whole file into memory
		var contents, err = ioutil.ReadFile(filename)

		if err != nil {
			fmt.Printf("Problem reading inputfile %s, error:%v\n", filename, err)
			return nil
		}

		var isBinary = utils.IsBinaryFormat(contents)

		// Binary files will not have "number of lines", but
		// will need to have the binary flag set for the report
		if isBinary {
			mutex.Lock()
			if !res.Extensions[ext].IsBinary {
				res.Extensions[ext].IsBinary = true
			}
			mutex.Unlock()
			return nil
		}

		var newline = utils.DetermineNewline(contents)
		var numberOfLines = int32(len(bytes.Split(contents, []byte(newline))))

		atomic.AddInt32(&(res.Extensions[ext].NumberOfLines), numberOfLines)
		atomic.AddInt32(&res.TotalNumberOfLines, numberOfLines)

		if numberOfBigFiles > 0 {
			mutex.Lock()
			res.BigFiles = append(res.BigFiles, result.FileSize{
				Name:  f.Name(),
				Size:  size,
				Lines: numberOfLines,
			})
			mutex.Unlock()
		}
	}

	return nil
}

// All returns the search result
func All(dir string, cfg config.Config) result.Result {
	startdir = dir
	res = result.InitResult(cfg.FileExtensions, cfg.Exclusions)

	var err = walk.Walk(startdir, forEachFileSystemEntry)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}

	return res
}
