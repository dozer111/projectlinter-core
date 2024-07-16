package substitute

import (
	"github.com/dozer111/projectlinter-core/rules"
)

type SubstitutePHPLibraryRule struct {
	*substituteLibraryRule
}

var _ rules.Rule = (*SubstitutePHPLibraryRule)(nil)

func NewSubstitutePHPLibraryRule(
	setName string,
	substituteRules []Library,
	dependencies []Dependency,
) *SubstitutePHPLibraryRule {
	return &SubstitutePHPLibraryRule{
		newSubstituteLibraryRule(setName, substituteRules, dependencies),
	}
}

// Validate validation compare exact match of the names in the projectlinter bump config and composer.json
func (r *SubstitutePHPLibraryRule) Validate() {
	dependenciesMap := make(dependencyMap, len(r.dependencies))
	for _, d := range r.dependencies {
		dependenciesMap[d.Name()] = d
	}

	for _, library := range r.substituteRules {
		libraryName := library.Name
		if dependenciesMap.Has(libraryName) {
			r.librariesToChange = append(r.librariesToChange, library)
		}
	}

	r.isPassed = len(r.librariesToChange) == 0
}
