package substitute

import (
	"fmt"
)

// substituteLibraryRule A mechanism for scheduled replacement of old libraries with new ones
type substituteLibraryRule struct {
	// projectlinter parsed configs
	substituteRules []Library
	// the libraries to substitute
	librariesToChange []Library
	dependencies      []Dependency
	setName           string
	isPassed          bool
}

type Dependency interface {
	Name() string
}

type dependencyMap map[string]Dependency

func (m dependencyMap) Has(name string) bool {
	_, ok := m[name]
	return ok
}

func newSubstituteLibraryRule(
	setName string,
	substituteRules []Library,
	dependencies []Dependency,
) *substituteLibraryRule {
	return &substituteLibraryRule{
		setName:         setName,
		substituteRules: substituteRules,
		dependencies:    dependencies,
	}
}

func (r *substituteLibraryRule) ID() string {
	return fmt.Sprintf("%s.substitute_library", r.setName)
}

func (r *substituteLibraryRule) Title() string {
	return "Substitute old libraries to new"
}

func (r *substituteLibraryRule) IsPassed() bool {
	return r.isPassed
}

func (r *substituteLibraryRule) FailedMessage() []string {
	return (&substituteLibraryPrinter{r.librariesToChange}).Print()
}
