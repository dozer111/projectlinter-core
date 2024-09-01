package bump

import (
	"github.com/Masterminds/semver/v3"
	"github.com/dozer111/projectlinter-core/rules"
)

type BumpJavascriptNPMLibraryRule struct {
	*bumpLibraryRule
}

var _ rules.Rule = (*BumpJavascriptNPMLibraryRule)(nil)

func NewBumpJavascriptNPMibraryRule(
	setName string,
	configs []Library,
	dependencies []Dependency,
) *BumpJavascriptNPMLibraryRule {
	return &BumpJavascriptNPMLibraryRule{
		newBumpLibraryRule(setName, configs, dependencies),
	}
}

// Validate validation compare exact match of the names in the projectlinter bump config and composer.json
func (r *BumpJavascriptNPMLibraryRule) Validate() {
	dependenciesMap := make(javascriptNPMDependencyMap, len(r.dependencies))

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
			// The situation when the library version in package.json and package-lock.json is different is possible
			if dependency.Version() != nil &&
				dependency.Version().LessThan(proposedVersion) {
				r.addLibrary(library, dependency)
			}
		}
	}

	r.isPassed = len(r.librariesToBump) == 0
}

type javascriptNPMDependencyMap map[string]Dependency

func (m javascriptNPMDependencyMap) Has(name string) bool {
	_, ok := m[name]
	return ok
}
