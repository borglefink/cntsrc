// Copyright 2017 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"os"
	"reflect"
	"testing"
)

// cleanupExclusions tests
func TestCleanupExclusions(t *testing.T) {
	var cfg = Config{
		Exclusions: []string{"\\dir", "\\bin\\", ".gitignore"},
	}
	var pss = string(os.PathSeparator)

	var actualResult = cfg.cleanupExclusions().Exclusions
	var expectedResult = []string{pss + "dir", pss + "bin" + pss, ".gitignore"}

	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

// getExecutableConfigName tests
func TestGetExecutableConfigName(t *testing.T) {
	var executableName = "cntsrc"

	var actualResult = getExecutableConfigName(executableName)
	var expectedResult = "cntsrc.config"

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestGetExecutableConfigNameExe(t *testing.T) {
	var executableName = "cntsrc.exe"

	var actualResult = getExecutableConfigName(executableName)
	var expectedResult = "cntsrc.config"

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestGetExecutableConfigNamePath(t *testing.T) {
	var path = "\\users\\user\\code\\bin\\"
	var executableName = "cntsrc"

	var actualResult = getExecutableConfigName(path + executableName)
	var expectedResult = path + "cntsrc.config"

	if actualResult != expectedResult {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

/*


*/
