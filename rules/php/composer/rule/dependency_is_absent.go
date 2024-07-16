package rule

import (
	"fmt"
	"github.com/dozer111/projectlinter-core/printer"
	"github.com/dozer111/projectlinter-core/rules"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
)

type DependencyIsAbsentRule struct {
	dependencies   *composer_json.ComposerDependencies
	dependencyName string
	isPassed       bool

	// for FailedMessage
	dependency     *composer_json.ComposerDependency
	additionalInfo []string
}

var _ rules.Rule = (*DependencyIsAbsentRule)(nil)

func NewComposerDependencyIsAbsentRule(
	dependencies *composer_json.ComposerDependencies,
	dependencyName string,
	additionalInfo []string,
) *DependencyIsAbsentRule {
	return &DependencyIsAbsentRule{
		dependencies:   dependencies,
		dependencyName: dependencyName,
		additionalInfo: additionalInfo,
	}
}

func (r *DependencyIsAbsentRule) ID() string {
	return fmt.Sprintf("composer.dependency.%s.absent", r.dependencyName)
}

func (r *DependencyIsAbsentRule) Title() string {
	return fmt.Sprintf(`dependency "%s" is absent`, r.dependencyName)
}

func (r *DependencyIsAbsentRule) Validate() {
	if r.dependencies.Has(r.dependencyName) {
		r.dependency = r.dependencies.Get(r.dependencyName)
		r.isPassed = false
		return
	}

	r.isPassed = true
}

func (r *DependencyIsAbsentRule) IsPassed() bool {
	return r.isPassed
}

func (r *DependencyIsAbsentRule) FailedMessage() []string {
	message := printer.NewCodeHintPrinter(
		[]string{fmt.Sprintf(`"%s": "%s",`, r.dependencyName, r.dependency.Constraint())},
		nil,
		nil,
		nil,
	).Print()

	if len(r.additionalInfo) > 0 {
		message = append(message, "", "Additional info:", "")
		for _, additionalMessage := range r.additionalInfo {
			message = append(message, fmt.Sprintf("  %s", additionalMessage))
		}
	}

	return message
}
