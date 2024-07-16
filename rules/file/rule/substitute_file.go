package file

import (
	"fmt"
	"github.com/dozer111/projectlinter-core/rules"
	"github.com/huandu/xstrings"
	"os"
)

// SubstituteFileRule checks whether the specified file exists, and if so, offers to replace it
type SubstituteFileRule struct {
	files          []os.DirEntry
	oldFileName    string
	newFileName    string
	linkToNewFile  string
	additionalInfo []string
	isPassed       bool
}

var _ rules.Rule = (*SubstituteFileRule)(nil)

func NewSubstituteFileRule(
	files []os.DirEntry,
	oldFileName string,
	newFileName string,
	linkToNewFile string,
	additionalInfo []string,
) *SubstituteFileRule {
	rule := &SubstituteFileRule{
		files:          files,
		oldFileName:    oldFileName,
		newFileName:    newFileName,
		linkToNewFile:  linkToNewFile,
		additionalInfo: additionalInfo,
	}

	return rule
}

func (r *SubstituteFileRule) ID() string {
	return fmt.Sprintf("file.substitute.%s", xstrings.ToSnakeCase(r.oldFileName))
}

func (r *SubstituteFileRule) Title() string {
	return fmt.Sprintf(`substitute file "%s"`, r.oldFileName)
}

func (r *SubstituteFileRule) Validate() {
	for _, fileInfo := range r.files {
		if !fileInfo.IsDir() && fileInfo.Name() == r.oldFileName {
			r.isPassed = false
			return
		}
	}

	r.isPassed = true
}

func (r *SubstituteFileRule) IsPassed() bool {
	return r.isPassed
}

func (r *SubstituteFileRule) FailedMessage() []string {
	result := []string{
		fmt.Sprintf("For some reason file \"%s\" is deprecated", r.oldFileName),
		fmt.Sprintf("Delete it, and copy new file \"%s\" from %s", r.newFileName, r.linkToNewFile),
	}

	if len(r.additionalInfo) > 0 {
		result = append(result, "", "Additional info:")
		result = append(result, r.additionalInfo...)
	}

	return result
}
