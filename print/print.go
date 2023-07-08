// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package print

import (
	"fmt"
	"sort"
	"strings"

	"cntsrc/result"
	"cntsrc/utils"
)

const (
	// formatString consists of "filetype", "#files", "#lines", "line%", "size", "size%"
	formatString       = "%-11s %10s %12s %6s %13s %6s\n"
	formatStringLength = 11 + 1 + 10 + 1 + 12 + 1 + 6 + 1 + 13 + 1 + 6
)

var (
	thousandsSeparator = ' '
)

// DebugHeader prints the debug indicator header
func DebugHeader() {
	fmt.Printf("\nShows included/excluded directories/files.\n")
	fmt.Printf("--------------------------------------------\n")
}

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

// printBigFiles
func printBigFiles(res result.Result) {
	sort.Sort(res.BigFiles)
	var label = fmt.Sprintf("The %d largest files are:", res.NumberOfBigFiles)
	fmt.Printf("\n\n%-56s #lines\n", label)
	fmt.Printf("---------------------------------------------------------------\n")

	for i := 0; i < res.NumberOfBigFiles; i++ {
		if i < len(res.BigFiles) {
			if i%10 == 0 {
				fmt.Println()
			}
			fmt.Printf("%-56s %6d\n", res.BigFiles[i].Name, res.BigFiles[i].Lines)
		}
	}
}

// getKeys
func getKeys(extensions map[string]*result.ExtensionEntry) ([]string, []string) {
	var keys []string
	var binaryKeys []string

	for k, v := range extensions {
		if v.NumberOfFiles > 0 {
			if v.IsBinary {
				binaryKeys = append(binaryKeys, k)
			} else {
				keys = append(keys, k)
			}
		}
	}

	sort.Strings(keys)
	sort.Strings(binaryKeys)

	return keys, binaryKeys
}

// printHeader
func printHeader(startdir string) {
	fmt.Printf("\nDirectory processed:\n")
	fmt.Printf("%v\n", startdir)
	fmt.Printf("%s\n", strings.Repeat("-", formatStringLength))
	fmt.Printf(formatString, "filetype", "#files", "#lines", "line%", "size", "size%")
	fmt.Printf("%s\n", strings.Repeat("-", formatStringLength))
}

// printEntry
func printEntry(entry *result.ExtensionEntry, totalNumberOfLines int32, totalSize int64) {
	var numberOfLinesString = ""
	var percentageString = ""

	if !entry.IsBinary {
		numberOfLinesString = utils.Int64ToString(int64(entry.NumberOfLines), thousandsSeparator)

		// Show percentage
		if totalNumberOfLines > 0 {
			var percentage = float64(entry.NumberOfLines) * float64(100) / float64(totalNumberOfLines)
			percentageString = fmt.Sprintf("%.1f", utils.Round(percentage, 1))
		}
	}

	var sizePercentage = float64(entry.Filesize) * float64(100) / float64(totalSize)
	var sizePercentageString = fmt.Sprintf("%.1f", utils.Round(sizePercentage, 1))

	fmt.Printf(
		formatString,
		entry.ExtensionName,
		utils.Int64ToString(int64(entry.NumberOfFiles), thousandsSeparator),
		numberOfLinesString,
		percentageString,
		utils.Int64ToString(int64(entry.Filesize), thousandsSeparator),
		sizePercentageString,
	)
}

// printFooter
func printFooter(res result.Result) {
	// Show footer
	if res.TotalNumberOfFiles == 0 {
		fmt.Printf("No files found.\n\nCheck given directory,or maybe \ncheck extensions in config file.\n")
		fmt.Printf("%s\n", strings.Repeat("-", formatStringLength))
		return
	}

	fmt.Printf("%s\n", strings.Repeat("-", formatStringLength))
	fmt.Printf(
		formatString,
		"Total:",
		utils.Int64ToString(int64(res.TotalNumberOfFiles), thousandsSeparator),
		utils.Int64ToString(int64(res.TotalNumberOfLines), thousandsSeparator),
		"100.0",
		utils.Int64ToString(int64(res.TotalSize), thousandsSeparator),
		"100.0",
	)
}
