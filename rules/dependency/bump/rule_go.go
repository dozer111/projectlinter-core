package bump

import (
	"github.com/dozer111/projectlinter-core/rules"
	"regexp"

	"github.com/Masterminds/semver/v3"
)

type BumpGOLibraryRule struct {
	*bumpLibraryRule
}

var _ rules.Rule = (*BumpGOLibraryRule)(nil)

func NewBumpGOLibraryRule(
	setName string,
	configs []Library,
	dependencies []Dependency,
) *BumpGOLibraryRule {
	return &BumpGOLibraryRule{
		newBumpLibraryRule(setName, configs, dependencies),
	}
}

// Validate validation takes place by the so-called "pure name"
// The point is that the GO package name can have a version suffix, for example "your_git.com/libraryName/v2"
//
// This value can be both in projectlinter bump config or in go.mod
// So, we cannot compare versions "head to head" as in PHP
//
// For comparison, we first reduce all names to pure (without suffix)
func (r *BumpGOLibraryRule) Validate() {
	goLibraryNameWithVersion := regexp.MustCompile(`.*/v\d+$`)
	goLibraryNameWithoutVersion := regexp.MustCompile(`(.*)/v\d+$`)

	cleanDependencies := make(goDependencyMap, len(r.dependencies))

	for _, d := range r.dependencies {
		dName := d.Name()
		if goLibraryNameWithVersion.MatchString(d.Name()) {
			dName = goLibraryNameWithoutVersion.FindStringSubmatch(dName)[1]
		}
		if len(cleanDependencies[dName]) == 0 {
			cleanDependencies[dName] = make([]Dependency, 0, 2)
		}

		cleanDependencies[dName] = append(cleanDependencies[dName], d)
	}

	for _, library := range r.configs {
		cleanLibraryName := library.Name
		if goLibraryNameWithVersion.MatchString(cleanLibraryName) {
			cleanLibraryName = goLibraryNameWithoutVersion.FindStringSubmatch(cleanLibraryName)[1]
		}

		if cleanDependencies.Has(cleanLibraryName) {
			version := library.Version
			proposedVersion, _ := semver.NewVersion(version)

			for _, dependency := range cleanDependencies[cleanLibraryName] {
				// version is the ONE source of truth
				if dependency.Version() != nil &&
					dependency.Version().LessThan(proposedVersion) {
					r.addLibrary(library, dependency)
					break
				}
			}
		}
	}

	r.isPassed = len(r.librariesToBump) == 0
}

type goDependencyMap map[string][]Dependency

func (m goDependencyMap) Has(name string) bool {
	_, ok := m[name]
	return ok
}
