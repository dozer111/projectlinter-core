package rule

import (
	"fmt"
	"github.com/dozer111/projectlinter-core/rules"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	utilSet "github.com/dozer111/projectlinter-core/util/set"
	"regexp"
)

// DependenciesConstrainsAreValidRule
// TODO add some output that when we have a library in exception and it is the correct version - show the user that this library can be removed from exceptions
type DependenciesConstrainsAreValidRule struct {
	dependencies *composer_json.ComposerDependencies
	// exceptions list of libraries that brake this rule
	// The idea is that some dependencies are known to never have a "correct" dependency
	// For example: "ext-json": "*",
	//
	// Therefore, we deliberately exclude them
	exceptions *utilSet.Set[string]

	depsWithWrongConstraint *composer_json.ComposerDependencies
	isPassed                bool
}

var _ rules.Rule = (*DependenciesConstrainsAreValidRule)(nil)

func NewDependenciesConstrainsAreValidRule(
	deps *composer_json.ComposerDependencies,
	exceptions []string,
) *DependenciesConstrainsAreValidRule {
	return &DependenciesConstrainsAreValidRule{
		dependencies:            deps,
		exceptions:              utilSet.NewSet[string](exceptions...),
		depsWithWrongConstraint: composer_json.NewComposerDependencies(),
	}
}

func (r *DependenciesConstrainsAreValidRule) ID() string {
	return "composer.dependencies_constraints_valid"
}

func (r *DependenciesConstrainsAreValidRule) Title() string {
	return `dependencies has correct constraint`
}

func (r *DependenciesConstrainsAreValidRule) Validate() {
	r1 := regexp.MustCompile(`^\^\d+$`)           // ^3
	r2 := regexp.MustCompile(`^\^\d+\.\d+$`)      // ^3.2
	r3 := regexp.MustCompile(`^\^\d+\.\d+\.\d+$`) // ^3.2.15

	for _, d := range r.dependencies.All() {
		if r1.MatchString(d.Constraint()) ||
			r2.MatchString(d.Constraint()) ||
			r3.MatchString(d.Constraint()) {
			continue
		}

		if r.exceptions.Has(d.Name()) {
			continue
		}

		r.depsWithWrongConstraint.Add(d)
	}

	r.isPassed = r.depsWithWrongConstraint.Count() == 0
}

func (r *DependenciesConstrainsAreValidRule) IsPassed() bool {
	return r.isPassed
}

func (r *DependenciesConstrainsAreValidRule) FailedMessage() []string {
	result := []string{
		`Dependency constraint(except "ext-") must match one of patterns: ^d+.d+(^8.1) or ^d+.d+.d+(^8.1.2)`,
		"- ^d+$         ^8",
		"- ^d+.d+$      ^8.1",
		"- ^d+.d+.d+$   ^8.1.2",
		"",
		`These libraries does not match it:`,
	}

	for _, r := range r.depsWithWrongConstraint.All() {
		result = append(result, fmt.Sprintf("	\"%s\": \"%s\"", r.Name(), r.Constraint()))
	}

	return result
}
