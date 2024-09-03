package substitute

import (
	"regexp"

	"github.com/dozer111/projectlinter-core/rules"
	"github.com/dozer111/projectlinter-core/rules/golang/gomod/config"
)

type SubstituteGOLibraryRule struct {
	*substituteLibraryRule
}

var _ rules.Rule = (*SubstituteGOLibraryRule)(nil)

func NewSubstituteGOLibraryRule(
	setName string,
	substituteRules []Library,
	dependencies []Dependency,
) *SubstituteGOLibraryRule {
	return &SubstituteGOLibraryRule{
		newSubstituteLibraryRule(setName, substituteRules, dependencies),
	}
}

// Validate validation takes place by the so-called "pure name"
// The point is that the GO package name can have a version suffix, for example "your_git.com/libraryName/v2"
//
// This value can be both in projectlinter bump config or in go.mod
// So, we cannot compare versions "head to head" as in PHP
//
// For comparison, we first reduce all names to pure (without suffix)
func (r *SubstituteGOLibraryRule) Validate() {
	goLibraryNameWithVersion := regexp.MustCompile(`.*/v\d+$`)
	goLibraryNameWithoutVersion := regexp.MustCompile(`(.*)/v\d+$`)

	cleanDependencies := make(dependencyMap, len(r.dependencies))

	for _, d := range r.dependencies {

		if d.(*config.GomodDependency).IsIndirect() {
			continue
		}

		dName := d.Name()
		if goLibraryNameWithVersion.MatchString(d.Name()) {
			dName = goLibraryNameWithoutVersion.FindStringSubmatch(dName)[1]
		}
		cleanDependencies[dName] = d
	}

	for _, library := range r.substituteRules {
		cleanLibraryName := library.Name
		if goLibraryNameWithVersion.MatchString(cleanLibraryName) {
			cleanLibraryName = goLibraryNameWithoutVersion.FindStringSubmatch(cleanLibraryName)[1]
		}

		if cleanDependencies.Has(cleanLibraryName) {
			r.librariesToChange = append(r.librariesToChange, library)
		}
	}

	r.isPassed = len(r.librariesToChange) == 0
}
