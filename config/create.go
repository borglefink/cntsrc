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
	var sc Config

	sc.FileExtensions = []string{".go", ".exe", ".png"}
	sc.Exclusions = []string{".git", filepath.Base(configFilename)}

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
