package rule

import (
	"fmt"
	"github.com/dozer111/projectlinter-core/printer"
	"github.com/dozer111/projectlinter-core/rules"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"github.com/huandu/xstrings"
)

// SpecialDependencyExistsRule - checks whether we have a special (which cannot be added by "composer install") specified dependency,
// and if not - recommends installing.
// There is also a similar ComposerDependencyExistsRule rule that checks dependencies that can be added by the composer
type SpecialDependencyExistsRule struct {
	isPassed        bool
	dependencies    *composer_json.ComposerDependencies
	dependencyName  string
	proposedVersion string
	// (for FailedMessage) - create more specified hint
	placedInRequireDev bool
}

var _ rules.Rule = (*SpecialDependencyExistsRule)(nil)

func NewSpecialDependencyExistsRule(
	dependencies *composer_json.ComposerDependencies,
	dependencyName,
	proposedVersion string,
	placedInRequireDev bool,
) *SpecialDependencyExistsRule {
	return &SpecialDependencyExistsRule{
		dependencies:       dependencies,
		dependencyName:     dependencyName,
		proposedVersion:    proposedVersion,
		placedInRequireDev: placedInRequireDev,
	}
}

func (r *SpecialDependencyExistsRule) ID() string {
	return fmt.Sprintf("composer.special_dependency.%s.exists", xstrings.ToSnakeCase(r.dependencyName))
}

func (r *SpecialDependencyExistsRule) Title() string {
	return fmt.Sprintf(`dependency "%s" exists`, r.dependencyName)
}

func (r *SpecialDependencyExistsRule) Validate() {
	r.isPassed = r.dependencies.Has(r.dependencyName)
}

func (r *SpecialDependencyExistsRule) IsPassed() bool {
	return r.isPassed
}

func (r *SpecialDependencyExistsRule) FailedMessage() []string {
	var beforeCode []string

	if r.placedInRequireDev {
		beforeCode = []string{`	"require-dev": {`}
	} else {
		beforeCode = []string{`	"require": {`}
	}

	return printer.NewCodeHintPrinter(
		beforeCode,
		nil,
		[]string{fmt.Sprintf(`		"%s": "%s",`, r.dependencyName, r.proposedVersion)},
		[]string{`		...`, `		...`},
	).Print()
}
