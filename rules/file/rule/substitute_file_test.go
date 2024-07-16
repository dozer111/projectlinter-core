package file_test

import (
	file "github.com/dozer111/projectlinter-core/rules/file/rule"
	utilTest "github.com/dozer111/projectlinter-core/util/test"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstituteFileRulePassWhileFileNameIsNewName(t *testing.T) {
	pathToService := utilTest.PathInProjectLinter("testdata")
	rootDirEntries, _ := os.ReadDir(pathToService)

	r := file.NewSubstituteFileRule(
		rootDirEntries,
		".php_cs.dist",
		".php-cs-fixer.dist.php",
		"",
		[]string{},
	)
	r.Validate()

	assert.True(t, r.IsPassed())
}

func TestSubstituteFileRuleFailWhileFileNameIsOldName(t *testing.T) {
	pathToService := utilTest.PathInProjectLinter("testdata")
	rootDirEntries, _ := os.ReadDir(pathToService)

	r := file.NewSubstituteFileRule(
		rootDirEntries,
		".php-cs-fixer.dist.php",
		".php-cs-fixer.dist",
		"",
		[]string{},
	)
	r.Validate()

	assert.False(t, r.IsPassed())
}
