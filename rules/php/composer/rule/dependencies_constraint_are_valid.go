package rule

import (
	"fmt"

	"github.com/dozer111/projectlinter-core/rules"

	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
)

type DependenciesConstrainsAreValidRule struct {
	dependencies   *composer_json.ComposerDependencies
	validationFunc func(string) bool

	depsWithWrongConstraint     *composer_json.ComposerDependencies
	additionalFailedMessageText []string
	isPassed                    bool
}

var _ rules.Rule = (*DependenciesConstrainsAreValidRule)(nil)

func NewDependenciesConstrainsAreValidRule(
	deps *composer_json.ComposerDependencies,
	validationFunc func(string) bool,
	additionalFailedMessageText []string,
) *DependenciesConstrainsAreValidRule {
	return &DependenciesConstrainsAreValidRule{
		dependencies:                deps,
		validationFunc:              validationFunc,
		additionalFailedMessageText: additionalFailedMessageText,
		depsWithWrongConstraint:     composer_json.NewComposerDependencies(),
	}
}

func (r *DependenciesConstrainsAreValidRule) ID() string {
	return "composer.dependencies_constraints_valid"
}

func (r *DependenciesConstrainsAreValidRule) Title() string {
	return `dependencies has correct constraint`
}

func (r *DependenciesConstrainsAreValidRule) Validate() {
	for _, d := range r.dependencies.All() {
		if !r.validationFunc(d.Constraint()) {
			r.depsWithWrongConstraint.Add(d)
		}
	}

	r.isPassed = r.depsWithWrongConstraint.Count() == 0
}

func (r *DependenciesConstrainsAreValidRule) IsPassed() bool {
	return r.isPassed
}

func (r *DependenciesConstrainsAreValidRule) FailedMessage() []string {
	result := r.additionalFailedMessageText

	result = append(
		result,
		"",
		"This library dependenies is wrong:",
	)

	for _, depWithWrongConstraint := range r.depsWithWrongConstraint.All() {
		result = append(
			result,
			fmt.Sprintf(
				"	\"%s\": \"%s\"",
				depWithWrongConstraint.Name(),
				depWithWrongConstraint.Constraint(),
			),
		)
	}

	return result
}
