// Copyright 2014-2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// createConfig creates the config, if it doesn't exist
func createConfig(configFilename string) Config {
	var sc = Config{
		FileExtensions: []string{".go", ".css", ".js", ".html", ".png", ".jpg"},
		Exclusions:     []string{".git", "\\bin\\", "\\node_modules\\", filepath.Base(configFilename)},
	}

	var jsonstring, err = json.MarshalIndent(&sc, "", "  ")
	if err != nil {
		fmt.Printf("Everything is wrong. json.Marshal(sc) didn't work; %s, error: %v\n", string(jsonstring), err)
	}

	err = ioutil.WriteFile(configFilename, jsonstring, 0666)
	if err != nil {
		fmt.Printf("Couldn't write the %s to disk, %v\n", configFilename, err)
	}

	return sc
}
