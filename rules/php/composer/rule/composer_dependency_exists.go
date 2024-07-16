package rule

import (
	"fmt"
	"github.com/dozer111/projectlinter-core/rules"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"github.com/huandu/xstrings"
)

// ComposerDependencyExistsRule checks whether we have the specified dependency, and if not - recommends installing it.
// Rule similar to SpecialDependencyExistsRule. And yet the SpecialDependencyExistsRule is issued separately,
// because there are dependencies for which we cannot do "composer install"(example: ext-json)
type ComposerDependencyExistsRule struct {
	isPassed       bool
	dependencies   *composer_json.ComposerDependencies
	dependencyName string

	// (for FailedMessage) - create more specified hint
	placedInRequireDev bool
	// usually the library is not added just like that. Plus, it was enough to add it. It would be nice to understand how to manage it
	// or some other work-specific related information.
	// Therefore, a more "verbose" description about how and why - must have
	readAlso []string
}

var _ rules.Rule = (*ComposerDependencyExistsRule)(nil)

func NewComposerDependencyExistsRule(
	dependencies *composer_json.ComposerDependencies,
	dependencyName string,
	placedInRequireDev bool,
	readAlso []string,
) *ComposerDependencyExistsRule {
	return &ComposerDependencyExistsRule{
		dependencies:       dependencies,
		dependencyName:     dependencyName,
		placedInRequireDev: placedInRequireDev,
		readAlso:           readAlso,
	}
}

func (r *ComposerDependencyExistsRule) ID() string {
	return fmt.Sprintf("composer.dependency.%s.exists", xstrings.ToSnakeCase(r.dependencyName))
}

func (r *ComposerDependencyExistsRule) Title() string {
	return fmt.Sprintf(`dependency "%s" exists`, r.dependencyName)
}

func (r *ComposerDependencyExistsRule) Validate() {
	r.isPassed = r.dependencies.Has(r.dependencyName)
}

func (r *ComposerDependencyExistsRule) IsPassed() bool {
	return r.isPassed
}

func (r *ComposerDependencyExistsRule) FailedMessage() []string {
	var result []string

	if r.placedInRequireDev {
		result = append(result, fmt.Sprintf("composer --dev require %s", r.dependencyName), "")
	} else {
		result = append(result, fmt.Sprintf("composer require %s", r.dependencyName), "")
	}

	if len(r.readAlso) > 0 {
		result = append(result, "Additional info:")
		for _, ra := range r.readAlso {
			result = append(result, fmt.Sprintf("  - %s", ra))
		}
	}

	return result
}
