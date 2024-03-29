// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// createConfig creates the config, if it doesn't exist
func createConfig(configFilename string) Config {
	var sc = Config{
		Inclusions: []string{".go", ".css", ".js", ".html", ".png", ".jpg"},
		Exclusions: []string{".git", "\\bin\\", "\\node_modules\\", filepath.Base(configFilename)},
	}

	var jsonstring, err = json.MarshalIndent(&sc, "", "  ")
	if err != nil {
		fmt.Printf("Help, help! json.Marshal(sc) didn't work. JSON: %s ERROR: %v\n", string(jsonstring), err)
	}

	err = os.WriteFile(configFilename, jsonstring, 0666)
	if err != nil {
		fmt.Printf("Couldn't write the %s to disk, %v\n", configFilename, err)
	}

	return sc
}
