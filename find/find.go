package find

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/MichaelTJones/walk"

	"cntsrc/config"
	"cntsrc/result"
	"cntsrc/utils"
)

var (
	res                  result.Result
	startdir             = ""
	debug                = false
	currentPathSeparator = string(os.PathSeparator)
)

// isExcluded
func isExcluded(filename string) bool {
	var fulldir, _ = filepath.Abs(filename)

	return utils.IsInString(fulldir, res.Exclusions)
}

// showDirectoriesOrFile
func showDirectoriesOrFile(isDir bool, filename string, excluded bool) {
	var status string

	if debug && isDir {
		if excluded {
			status = " EXCLUDED "
		} else {
			status = "          "
		}

		fmt.Printf("Directory %s%s\n", status, strings.Replace(filename, startdir+currentPathSeparator, "", 1))
	}

	if debug && !isDir {
		if excluded {
			status = " EXCLUDED "
		} else {
			status = "          "
		}

		fmt.Printf("File      %s%s\n", status, strings.Replace(filename, startdir+currentPathSeparator, "", 1))
	}
}

// forEachFileSystemEntry
func forEachFileSystemEntry(filename string, f os.FileInfo, err error) error {

	if f == nil {
		return nil
	}

	var excluded = isExcluded(filename)

	if debug {
		showDirectoriesOrFile(f.IsDir(), filename, excluded)
	}

	if f.IsDir() {
		return nil
	}

	if !excluded {
		// Extension for the entry we're looking at
		var ext = filepath.Ext(filename)

		// Is the extension one of the relevant ones?
		var _, willBeCounted = res.Extensions[ext]

		// If no, exit
		if !willBeCounted {
			return nil
		}

		res.Extensions[ext].NumberOfFiles++
		res.TotalNumberOfFiles++

		var size = f.Size()
		res.Extensions[ext].Filesize += size
		res.TotalSize += size

		// Slurp the whole file into memory
		var contents, err = ioutil.ReadFile(filename)

		if err != nil {
			fmt.Printf("Problem reading inputfile %s, error:%v\n", filename, err)
			return nil
		}

		var isBinary = utils.IsBinaryFormat(contents)

		// Binary files will not have "number of lines", but
		// will need to have the binary flag set for the report
		if isBinary && !res.Extensions[ext].IsBinary {
			res.Extensions[ext].IsBinary = true
			return nil
		}

		var newline = utils.DetermineNewline(contents)

		var numberOfLines = len(bytes.Split(contents, []byte(newline)))

		res.Extensions[ext].NumberOfLines += numberOfLines
		res.TotalNumberOfLines += numberOfLines
		res.BigFiles = append(res.BigFiles, result.FileSize{
			Name:  f.Name(),
			Size:  size,
			Lines: numberOfLines,
		})
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
