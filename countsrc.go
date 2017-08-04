// Copyright 2014-2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MichaelTJones/walk"

	"countsrc/config"
)

var (
	startdir      = *flag.String("c", "", "countsource configuration file")
	showDebug     = *flag.Bool("d", false, "show full status of which files and directories in path are excluded or included.")
	showBigFiles  = *flag.Int("big", 0, "show the x largest files")
	showHelp      = *flag.Bool("?", false, "this help information")
	config        Config
	countResult   Result
	exclusions    Exclusions
	pathSeparator = getPathSeparator()
	bigFiles      = make(fileSizes, 0)
)

// init
func init() {
	flag.Usage = usage
	flag.Parse()

	if showHelp {
		usage()
	}

	if len(startdir) > 0 {
		startdir = 
	}

	// Find given drirectory
	startdir = getDirectory(flag.Arg(0), ".")

	// Load config and prepare for parsing directory
	var configFilename = getConfigFileName(*suggestedConfigFilename)
	config = loadConfig(configFilename)
	exclusions = config.getExclusions()
	countResult = config.setupResult()
	printAnalyticsHeader()
}

// usage
func usage() {
	var executableName = filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "\nCOUNTSOURCE (C) Copyright 2014-2017 Erlend Johannessen\n")
	fmt.Fprintf(os.Stderr, "%s counts sourcecode lines for given directory and sub-directories.\n", executableName)
	fmt.Fprintf(os.Stderr, "\nUsage: %s [dirname] [-c fullpathtoconfigfile] \n", executableName)
	fmt.Fprintf(os.Stderr, "  dirname: Name of directory with source code to count lines for. Uses current directory if no directory given.\n")
	flag.PrintDefaults()
	os.Exit(0)
}

// main
func main() {
	// Processing the given directory
	var err = walk.Walk(startdir, forEachFileSystemEntry)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}

	// Show result
	printResult()
}
