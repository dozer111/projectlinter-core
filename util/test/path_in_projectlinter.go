package utilTest

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// PathInProjectLinter is the full path to the folder/file inside the projectlinter
//
// This method is used exclusively in tests so that you can easily get to a certain folder/file without additional complications
func PathInProjectLinter(relativePath string) string {
	return fmt.Sprintf("%s/%s", pathToProjectlinter(), relativePath)
}

func pathToProjectlinter() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(strings.TrimSuffix(b, "util/test/path_in_projectlinter.go"))
}
