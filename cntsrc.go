// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"cntsrc/config"
	"cntsrc/find"
	"cntsrc/print"
	"cntsrc/utils"
)

var (
	suggestedConfigFilename = flag.String("c", "", "countsource configuration file")
	showDebug               = flag.Bool("debug", false, "show full status of which files and directories in path are excluded or included.")
	showBigFiles            = flag.Int("big", 0, "show the x largest files")
	showHelp                = flag.Bool("h", false, "this help information")
	startdir                = "."
	cfg                     config.Config
)

// init
func init() {
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		usage()
	}

	find.SetDebug(*showDebug)
	startdir = utils.ResolveStartdir(flag.Arg(0), ".")
	cfg = config.LoadConfig(*suggestedConfigFilename)
}

// usage
func usage() {
	var executableName = filepath.Base(os.Args[0])
	fmt.Printf("\nCNTSRC (C) Copyright 2017 Erlend Johannessen\n")
	fmt.Printf("%s counts source-code lines for given directory and sub-directories.\n", executableName)
	fmt.Printf("\nUsage: %s [options] [dirname]  \n", executableName)
	fmt.Printf("  dirname: Name of directory with source code to count lines for. Uses current directory if no directory given.\n")
	flag.PrintDefaults()
}

// main
func main() {
	print.Result(
		startdir,
		find.All(startdir, cfg, *showBigFiles),
	)
}
