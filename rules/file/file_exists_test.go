package file_test

import (
	"github.com/dozer111/projectlinter-core/rules/file"
	utilTest "github.com/dozer111/projectlinter-core/util/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExistsSuccessCase(t *testing.T) {
	r := file.NewFileExistsRule(
		utilTest.PathInProjectLinter("testdata"),
		".gitkeep",
		"",
	)
	r.Validate()

	assert.True(t, r.IsPassed())
}

func TestFileExistsFailWhileFileIsAbsent(t *testing.T) {
	r := file.NewFileExistsRule(
		utilTest.PathInProjectLinter("testdata"),
		"somePHPFile.php",
		"",
	)
	r.Validate()

	assert.False(t, r.IsPassed())
}
