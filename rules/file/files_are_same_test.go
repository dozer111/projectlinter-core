package file_test

import (
	"embed"
	"github.com/dozer111/projectlinter-core/rules/file"
	utilTest "github.com/dozer111/projectlinter-core/util/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/.editorconfig2
var editorconfig2 embed.FS

//go:embed testdata/.editorconfig3
var editorconfig3 embed.FS

func TestFilesAreSame(t *testing.T) {
	t.Run("pass on same files", func(t *testing.T) {
		f, err := editorconfig2.Open("testdata/.editorconfig2")
		assert.NoError(t, err)

		r := file.NewFilesAreSameRule(
			utilTest.PathInProjectLinter("testdata/.editorconfig"),
			f,
			"",
		)
		r.Validate()

		assert.True(t, r.IsPassed())
	})

	t.Run("fail on different files", func(t *testing.T) {
		f, err := editorconfig3.Open("testdata/.editorconfig3")
		assert.NoError(t, err)

		r := file.NewFilesAreSameRule(
			utilTest.PathInProjectLinter("testdata/.editorconfig"),
			f,
			"",
		)
		r.Validate()

		assert.False(t, r.IsPassed())
	})
}
