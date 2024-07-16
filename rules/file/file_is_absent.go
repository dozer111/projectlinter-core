package file

import (
	"fmt"
	"github.com/1set/gut/yos"
	"github.com/dozer111/projectlinter-core/rules"
	"github.com/huandu/xstrings"
)

// FileIsAbsentRule checks if the specified file is missing
//
// Why exactly 2 fields absolutePathToProject, relativePathToFile and not one common, through pathProvider.PathInCaller()?
//
// Because it is much easier to show which file is missing in the error response (bin/data, not just data, for example)
type FileIsAbsentRule struct {
	absolutePathToProject string
	relativePathToFile    string
	isPassed              bool
}

var _ rules.Rule = (*FileIsAbsentRule)(nil)

func NewFileIsAbsentRule(
	dirWithFile string,
	filepath string,
) *FileIsAbsentRule {
	return &FileIsAbsentRule{
		absolutePathToProject: dirWithFile,
		relativePathToFile:    filepath,
	}
}

func (r *FileIsAbsentRule) ID() string {
	return fmt.Sprintf("file.absent.%s", xstrings.ToSnakeCase(r.relativePathToFile))
}

func (r *FileIsAbsentRule) Title() string {
	return fmt.Sprintf("file \"%s\" is absent", r.relativePathToFile)
}

func (r *FileIsAbsentRule) Validate() {
	r.isPassed = !yos.ExistFile(fmt.Sprintf("%s/%s", r.absolutePathToProject, r.relativePathToFile))
}

func (r *FileIsAbsentRule) IsPassed() bool {
	return r.isPassed
}

func (r *FileIsAbsentRule) FailedMessage() []string {
	return []string{
		fmt.Sprintf("Delete file \"%s\"", r.relativePathToFile),
	}
}
