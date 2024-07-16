package file

import (
	"bytes"
	"fmt"
	"github.com/1set/gut/yos"
	"github.com/dozer111/projectlinter-core/rules"
	"github.com/huandu/xstrings"
	"io"
	"io/fs"
	"os"
	"strings"
)

type FilesAreSameRule struct {
	absolutePathToFile1 string
	embeddedFile        fs.File

	validationError error

	fileName     string
	linkOnSource string

	isPassed bool
}

var _ rules.Rule = (*FilesAreSameRule)(nil)

func NewFilesAreSameRule(
	absolutePathToFileInCaller string,
	embeddedFile fs.File,
	linkOnSource string,
) *FilesAreSameRule {
	path := strings.Split(absolutePathToFileInCaller, "/")
	fileName := path[len(path)-1]

	return &FilesAreSameRule{
		absolutePathToFile1: absolutePathToFileInCaller,
		embeddedFile:        embeddedFile,
		fileName:            fileName,
		linkOnSource:        linkOnSource,
	}
}

func (r *FilesAreSameRule) ID() string {
	return fmt.Sprintf("file.correct.%s", xstrings.ToSnakeCase(r.fileName))
}

func (r *FilesAreSameRule) Title() string {
	return fmt.Sprintf("file \"%s\" is correct", r.fileName)
}

func (r *FilesAreSameRule) Validate() {
	same, err := r.filesAreSame(r.absolutePathToFile1, r.embeddedFile)
	if err != nil {
		r.validationError = err
		r.isPassed = false
		return
	}

	r.isPassed = same
}

func (r *FilesAreSameRule) IsPassed() bool {
	return r.isPassed
}

func (r *FilesAreSameRule) FailedMessage() []string {
	result := []string{
		fmt.Sprintf("File %s is not equal to %s", r.fileName, r.linkOnSource),
	}

	if r.validationError != nil {
		result = append(result, fmt.Sprintf("Validation failed with error: %s", r.validationError))
	}

	return result
}

// filesAreSame https://stackoverflow.com/a/30038571/9500254
// Implemented with go:embed(because the origin file is placing on your side and you have access to it)
func (r *FilesAreSameRule) filesAreSame(pathToFile1 string, file2 fs.File) (bool, error) {
	// step1 check files exists
	if !yos.ExistFile(pathToFile1) {
		return false, fmt.Errorf("file %s does not exists", pathToFile1)
	}

	// step2 check files has same value
	{
		f1, err := os.Stat(pathToFile1)
		if err != nil {
			return false, fmt.Errorf("fail on os.Stat %s: %w", f1, err)
		}

		f2, err := file2.Stat()
		if err != nil {
			return false, fmt.Errorf("cannot read stat of projectlinter embedded file %s: %w", f1, err)
		}

		if f1.Size() != f2.Size() {
			return false, nil
		}
	}

	f1, err := os.Open(pathToFile1)
	if err != nil {
		return false, fmt.Errorf("cannot open file %s: %w", pathToFile1, err)
	}

	chunckSize := 64000

	b1 := make([]byte, chunckSize)
	_, err1 := f1.Read(b1)

	b2 := make([]byte, chunckSize)
	_, err2 := file2.Read(b2)

	if err1 != nil || err2 != nil {
		if err1 == io.EOF && err2 == io.EOF {
			return true, nil
		} else {
			if err1 != nil {
				return false, fmt.Errorf("failure on compare files, err1: %w", err1)
			}

			return false, fmt.Errorf("failure on compare files, err2: %w", err2)
		}
	}

	return bytes.Equal(b1, b2), nil
}
