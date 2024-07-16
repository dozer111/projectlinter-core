package file_test

import (
	"github.com/dozer111/projectlinter-core/rules/file"
	utilTest "github.com/dozer111/projectlinter-core/util/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenameFileRulePassWhileFileNameIsNewName(t *testing.T) {
	pathToService := utilTest.PathInProjectLinter("testdata")

	r := file.NewRenameFileRule(
		pathToService,
		"taglist.yaml",
		"tags.yaml",
	)
	r.Validate()

	assert.True(t, r.IsPassed())
}

func TestRenameFileRuleFailWhileFileNameIsOldName(t *testing.T) {
	pathToService := utilTest.PathInProjectLinter("testdata")

	r := file.NewRenameFileRule(
		pathToService,
		"tags.yaml",
		"tags-bundle.yaml",
	)
	r.Validate()

	assert.False(t, r.IsPassed())
}
