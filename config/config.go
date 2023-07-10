// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"cntsrc/utils"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	defaultConfigFileName = "cntsrc.config"
)

// Config contains the programs config, read from file
type Config struct {
	FileExtensions []string `json:"FileExtensions"`
	Exclusions     []string `json:"Exclusions"`
}

// cleanupExclusions
func (sc Config) cleanupExclusions() Config {
	var ps = string(os.PathSeparator)
	for i := range sc.Exclusions {
		sc.Exclusions[i] = strings.Replace(sc.Exclusions[i], "\\", ps, -1)
	}

	return sc
}

// getExecutableConfigName
func getExecutableConfigName(executableName string) string {
	var executableConfigName = executableName

	// if Windows platform, remove .exe
	if strings.Index(executableConfigName, ".exe") > 0 {
		executableConfigName = strings.Replace(executableConfigName, ".exe", "", 1)
	}

	return executableConfigName + ".config"

}

// stripAllComments
func stripAllComments(jsonstring []byte) []byte {
	var re = regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	return re.ReplaceAll(jsonstring, nil)
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
	var executableName = utils.GetExecutableName()
	var executableConfigName = getExecutableConfigName(executableName)
	if _, err := os.Stat(executableConfigName); err == nil {
		return executableConfigName
	}

	// if no other config file exist,
	// return config file for the current directory,
	// whether it exists or not
	return defaultConfigFileName
}

// LoadConfig loads the config from file
func LoadConfig(suggestedConfigFilename string) Config {
	var configFilename = resolveConfigFileName(suggestedConfigFilename)

	// Read whole the file. If not exist, create it.
	var jsonstring, err = os.ReadFile(configFilename)
	if err != nil {
		return createConfig(defaultConfigFileName).cleanupExclusions()
	}

	// Strip comments from config file
	var newJsonstring = stripAllComments(jsonstring)

	// Create config to be returned
	var c Config
	err = json.Unmarshal(newJsonstring, &c)

	if err != nil {
		fmt.Printf("Could not read the config, %v\n", err)
		return Config{}
	}

	return c.cleanupExclusions()
}
