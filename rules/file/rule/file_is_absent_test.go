package file_test

import (
	"fmt"
	"github.com/dozer111/projectlinter-core/rules/file/rule"
	utilTest "github.com/dozer111/projectlinter-core/util/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileIsAbsentRuleSuccessCase(t *testing.T) {
	r := file.NewFileIsAbsentRule(
		utilTest.PathInProjectLinter("testdata"),
		".gitignore",
	)
	r.Validate()

	fmt.Println(r.ID())
	assert.True(t, r.IsPassed())
}

func TestFileIsAbsentRuleFailWhileFileExists(t *testing.T) {
	r := file.NewFileIsAbsentRule(
		utilTest.PathInProjectLinter("testdata"),
		".gitkeep",
	)
	r.Validate()

	assert.False(t, r.IsPassed())
}
