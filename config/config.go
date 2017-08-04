// Copyright 2014-2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kardianos/osext"
)

const (
	defaultConfigFileName = "countsource.config"
)

// Config contains the programs config, read from file
type Config struct {
	Extensions         []string
	BinaryExtensions   []string
	ExcludeDirectories []string
	ExcludeFiles       []string
}

// setupExclusions
func (sc Config) getExclusions() Exclusions {
	var exclusions = Exclusions{sc.ExcludeDirectories, sc.ExcludeFiles}

	// Initialise ExcludeDirectories with searchable criteria
	for index := range exclusions.ExcludeDirectories {
		exclusions.ExcludeDirectories[index] = pathSeparator + exclusions.ExcludeDirectories[index] + pathSeparator
	}

	return exclusions
}

// setupResult
func (sc Config) setupResult() Result {
	var r = Result{Extensions: make(map[string]*ResultEntry)}

	// Add extensions
	for _, k := range sc.Extensions {
		r.Extensions[k] = &ResultEntry{k, false, 0, 0, 0}
	}

	// Add binary extensions
	for _, k := range sc.BinaryExtensions {
		r.Extensions[k] = &ResultEntry{k, true, 0, 0, 0}
	}

	return r
}

func getExecutableName() string {
	filename, err := osext.Executable()

	if err != nil {
		filename, _ = filepath.Abs(os.Args[0])
	}

	return filename
}

// getConfigFileName
func getConfigFileName(suggestedConfigFilename string) string {
	var fullFilePath = defaultConfigFileName

	if len(suggestedConfigFilename) > 0 {
		if _, err := os.Stat(suggestedConfigFilename); err == nil {
			fullFilePath = suggestedConfigFilename
			return fullFilePath
		}
	}

	if _, err := os.Stat(fullFilePath); os.IsNotExist(err) {
		fullFilePath = getExecutableName()

		if strings.Index(fullFilePath, ".exe") > 0 {
			fullFilePath = strings.Replace(fullFilePath, ".exe", "", 1)
		}

		fullFilePath = fullFilePath + ".config"
	}

	return fullFilePath
}

// loadConfig loads the config from file
func loadConfig(configFilename string) Config {
	var c Config

	// Read whole the file
	var jsonstring, err = ioutil.ReadFile(configFilename)
	if err != nil {
		c = createConfig(configFilename)
		return c
	}

	// Strip comments from config file
	var re = regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	var newJsonstring = re.ReplaceAll(jsonstring, nil)

	// Create config to be returned
	err = json.Unmarshal(newJsonstring, &c)

	if err != nil {
		fmt.Printf("Could not read the config, %v\n", err)
		c = Config{}
	}

	return c
}

// createConfig creates the config, if it doesn't exist
func createConfig(configFilename string) Config {
	var sc Config

	sc.Extensions = []string{".go"}
	sc.BinaryExtensions = []string{".exe", ".png"}
	sc.ExcludeDirectories = []string{".git"}
	sc.ExcludeFiles = []string{filepath.Base(configFilename)}

	var jsonstring, err = json.MarshalIndent(&sc, "", "  ")
	if err != nil {
		fmt.Printf("json.Marshal(sc), %s %v\n", string(jsonstring), err)
	}

	err = ioutil.WriteFile(configFilename, jsonstring, 0666)
	if err != nil {
		fmt.Printf("ioutil.WriteFile, %v\n", err)
	}

	return sc
}
