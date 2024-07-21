package platform

import (
	"github.com/dozer111/projectlinter-core/printer"
	"github.com/dozer111/projectlinter-core/rules"
)

// OnlySpecifiedPlatformExistsRule check that "config.platform" contains ONLY specified platforms
type OnlySpecifiedPlatformExistsRule struct {
	expectedValues map[string]string
	value          map[string]string
	isPassed       bool
	isValidated    bool
}

var _ rules.Rule = (*OnlySpecifiedPlatformExistsRule)(nil)

func NewOnlySpecifiedPlatformExistsRule(
	platformsInConfig map[string]string,
	expectedValues map[string]string,
) *OnlySpecifiedPlatformExistsRule {
	return &OnlySpecifiedPlatformExistsRule{
		value:          platformsInConfig,
		expectedValues: expectedValues,
	}
}

func (r *OnlySpecifiedPlatformExistsRule) ID() string {
	return "composer.config.platform.only_specified.exists"
}

func (r *OnlySpecifiedPlatformExistsRule) Title() string {
	return `"config.platform" has only specified values`
}

func (r *OnlySpecifiedPlatformExistsRule) Validate() {
	value := r.value

	if len(value) > len(r.expectedValues) {
		r.isPassed = false
		return
	}

	for platformName := range r.expectedValues {
		if _, ok := (value)[platformName]; !ok {
			r.isPassed = false
			return
		}
	}

	r.isPassed = true
}

func (r *OnlySpecifiedPlatformExistsRule) IsPassed() bool {
	return r.isPassed
}

func (r *OnlySpecifiedPlatformExistsRule) FailedMessage() []string {
	return printer.NewCodeHintPrinter(
		nil,
		platformPrettyJson(r.value),
		platformPrettyJson(r.expectedValues),
		nil,
	).Print()
}
