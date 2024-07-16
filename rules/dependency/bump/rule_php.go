package bump

import (
	"github.com/Masterminds/semver/v3"
	"github.com/dozer111/projectlinter-core/rules"
)

type BumpPHPLibraryRule struct {
	*bumpLibraryRule
}

var _ rules.Rule = (*BumpPHPLibraryRule)(nil)

func NewBumpPHPLibraryRule(
	setName string,
	configs []Library,
	dependencies []Dependency,
) *BumpPHPLibraryRule {
	return &BumpPHPLibraryRule{
		newBumpLibraryRule(setName, configs, dependencies),
	}
}

// Validate validation compare exact match of the names in the projectlinter bump config and composer.json
func (r *BumpPHPLibraryRule) Validate() {
	dependenciesMap := make(phpDependencyMap, len(r.dependencies))

	for _, d := range r.dependencies {
		dependenciesMap[d.Name()] = d
	}

	for _, library := range r.configs {
		version := library.Version
		libraryName := library.Name
		proposedVersion, _ := semver.NewVersion(version)

		if dependenciesMap.Has(libraryName) {
			dependency := dependenciesMap[libraryName]
			// version is the ONE source of truth
			if dependency.Version() != nil &&
				dependency.Version().LessThan(proposedVersion) {
				r.addLibrary(library, dependency)
			}
		}
	}

	r.isPassed = len(r.librariesToBump) == 0
}

type phpDependencyMap map[string]Dependency

func (m phpDependencyMap) Has(name string) bool {
	_, ok := m[name]
	return ok
}
