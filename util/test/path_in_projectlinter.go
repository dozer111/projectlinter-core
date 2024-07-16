package utilTest

import (
	"fmt"
	"os"
	"path/filepath"
)

// PathInProjectLinter is the full path to the folder/file inside the projectlinter
//
// This method is used exclusively in tests so that you can easily get to a certain folder/file without additional complications
func PathInProjectLinter(relativePath string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("cannot find caller dir:" + err.Error())
	}

	newPath := filepath.Join(currentDir, relativePath)
	return newPath
}
