// Copyright 2014-2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"countsrc/utils"
)

const (
	defaultConfigFileName = "countsrc.config"
)

var (
	pathSeparator = ""
)

// Config contains the programs config, read from file
type Config struct {
	FileExtensions []string
	Exclusions     []string
}

// cleanupExclusions
func (sc Config) cleanupExclusions() Config {
	var ps = string(os.PathSeparator)
	for i := range sc.Exclusions {
		sc.Exclusions[i] = strings.Replace(sc.Exclusions[i], "\\", ps, -1)
	}

	return sc
}

// resolveConfigFileName
func resolveConfigFileName(suggestedConfigFilename string) string {
	// config filename from cli parameters
	if len(suggestedConfigFilename) > 0 {
		if _, err := os.Stat(suggestedConfigFilename); err == nil {
			return suggestedConfigFilename
		}
	}

	// config filename from current directory
	if _, err := os.Stat(defaultConfigFileName); err == nil {
		return defaultConfigFileName
	}

	// config filename from the executable
	var executableConfigName = utils.GetExecutableName()

	// if Windows platform, remove .exe
	if strings.Index(executableConfigName, ".exe") > 0 {
		executableConfigName = strings.Replace(executableConfigName, ".exe", "", 1)
	}

	executableConfigName = executableConfigName + ".config"

	if _, err := os.Stat(executableConfigName); err == nil {
		return executableConfigName
	}

	// if no other config file exist,
	// return config file for the current directory
	return defaultConfigFileName
}

// LoadConfig loads the config from file
func LoadConfig(suggestedConfigFilename string) Config {
	var configFilename = resolveConfigFileName(suggestedConfigFilename)

	// Read whole the file. If not exist, create it.
	var jsonstring, err = ioutil.ReadFile(configFilename)
	if err != nil {
		return createConfig(configFilename).cleanupExclusions()
	}

	// Strip comments from config file
	var re = regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	var newJsonstring = re.ReplaceAll(jsonstring, nil)

	// Create config to be returned
	var c Config
	err = json.Unmarshal(newJsonstring, &c)

	if err != nil {
		fmt.Printf("Could not read the config, %v\n", err)
		return Config{}
	}

	return c.cleanupExclusions()
}
