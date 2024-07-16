package file

import (
	"fmt"
	"github.com/dozer111/projectlinter-core/rules"
	"github.com/huandu/xstrings"

	"github.com/1set/gut/yos"
)

// RenameFileRule checks whether the specified file exists, and if so, offers to rename it
type RenameFileRule struct {
	absolutePathToProject string

	oldFilePath string
	newFilePath string
	isPassed    bool
}

var _ rules.Rule = (*RenameFileRule)(nil)

func NewRenameFileRule(
	absolutePathToFile,
	oldFilePath,
	newFilePath string,
) *RenameFileRule {
	return &RenameFileRule{
		absolutePathToProject: absolutePathToFile,
		oldFilePath:           oldFilePath,
		newFilePath:           newFilePath,
	}
}

func (r *RenameFileRule) ID() string {
	return fmt.Sprintf("file.rename.%s", xstrings.ToSnakeCase(r.oldFilePath))
}

func (r *RenameFileRule) Title() string {
	return fmt.Sprintf(`rename file %s`, r.oldFilePath)
}

func (r *RenameFileRule) Validate() {
	if yos.ExistFile(fmt.Sprintf("%s/%s", r.absolutePathToProject, r.oldFilePath)) {
		r.isPassed = false
		return
	}

	r.isPassed = true
}

func (r *RenameFileRule) IsPassed() bool {
	return r.isPassed
}

func (r *RenameFileRule) FailedMessage() []string {
	return []string{
		fmt.Sprintf(`Rename file: %s => %s`, r.oldFilePath, r.newFilePath),
	}
}
