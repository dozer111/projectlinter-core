package bump

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
)

// bumpLibraryRule Планове безболісне підняття бібліотек.
type bumpLibraryRule struct {
	configs      []Library
	dependencies []Dependency

	librariesToBump []*bumpDependencyConfig
	setName         string

	isPassed bool
}

type Dependency interface {
	Name() string
	Version() *semver.Version
	VersionRaw() string
}

func newBumpLibraryRule(
	setName string,
	configs []Library,
	dependencies []Dependency,
) *bumpLibraryRule {
	return &bumpLibraryRule{
		setName:      setName,
		configs:      configs,
		dependencies: dependencies,
	}
}

func (r *bumpLibraryRule) ID() string {
	return fmt.Sprintf("%s.bump_library", r.setName)
}

func (r *bumpLibraryRule) Title() string {
	return "Bump libraries version"
}

func (r *bumpLibraryRule) addLibrary(cfg Library, dep Dependency) {
	r.librariesToBump = append(r.librariesToBump, &bumpDependencyConfig{
		Library:        cfg,
		CurrentVersion: dep.VersionRaw(),
	})
}

func (r *bumpLibraryRule) IsPassed() bool {
	return r.isPassed
}

func (r *bumpLibraryRule) FailedMessage() []string {
	return (&bumpDependencyPrinter{r.librariesToBump}).Print()
}
