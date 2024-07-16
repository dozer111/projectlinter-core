package file

import (
	"fmt"
	"github.com/1set/gut/yos"
	"github.com/dozer111/projectlinter-core/rules"
	"github.com/huandu/xstrings"
)

// FileExistsRule
//
// Why exactly 2 fields absolutePathToProject, relativePathToFile and not one common, through pathProvider.PathInCaller()?
//
// Because it is much easier to show which file is missing in the error response (bin/data, not just data, for example)
type FileExistsRule struct {
	absolutePathToProject string
	relativePathToFile    string

	linkOnSource string

	isPassed bool
}

var _ rules.Rule = (*FileExistsRule)(nil)

func NewFileExistsRule(
	dirWithFile,
	filepath,
	linkOnSource string,
) *FileExistsRule {
	return &FileExistsRule{
		absolutePathToProject: dirWithFile,
		relativePathToFile:    filepath,
		linkOnSource:          linkOnSource,
	}
}

func (r *FileExistsRule) ID() string {
	return fmt.Sprintf("file.exists.%s", xstrings.ToSnakeCase(r.relativePathToFile))
}

func (r *FileExistsRule) Title() string {
	return fmt.Sprintf("file \"%s\" exists", r.relativePathToFile)
}

func (r *FileExistsRule) Validate() {
	r.isPassed = yos.ExistFile(fmt.Sprintf("%s/%s", r.absolutePathToProject, r.relativePathToFile))
}

func (r *FileExistsRule) IsPassed() bool {
	return r.isPassed
}

func (r *FileExistsRule) FailedMessage() []string {
	return []string{
		fmt.Sprintf("Copy \"%s\" from %s", r.relativePathToFile, r.linkOnSource),
	}
}
