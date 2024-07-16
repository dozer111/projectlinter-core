package platform

import (
	"fmt"
	"github.com/dozer111/projectlinter-core/printer"
	"github.com/dozer111/projectlinter-core/rules"
	"strings"
)

// SpecifiedPlatformExistsRule Check that "config.platform" contains the specified platform
type SpecifiedPlatformExistsRule struct {
	expectedPlatform string
	possibleValue    string
	value            map[string]string
	isPassed         bool
}

var _ rules.Rule = (*SpecifiedPlatformExistsRule)(nil)

func NewSpecifiedPlatformExistsRule(
	platforms map[string]string,
	expectedPlatform,
	possibleValue string,
) *SpecifiedPlatformExistsRule {
	return &SpecifiedPlatformExistsRule{
		expectedPlatform: expectedPlatform,
		possibleValue:    possibleValue,
		value:            platforms,
	}
}

func (r *SpecifiedPlatformExistsRule) ID() string {
	return fmt.Sprintf("composer.config.platform.%s.exists", r.expectedPlatform)
}

func (r *SpecifiedPlatformExistsRule) Title() string {
	return fmt.Sprintf(`"config.platform.%s" exists`, r.expectedPlatform)
}

func (r *SpecifiedPlatformExistsRule) Validate() {
	_, phpExists := (r.value)[r.expectedPlatform]
	r.isPassed = phpExists
}

func (r *SpecifiedPlatformExistsRule) IsPassed() bool {
	return r.isPassed
}

func (r *SpecifiedPlatformExistsRule) FailedMessage() []string {
	currentPlatforms := make(map[string]string, len(r.value))
	for p, v := range r.value {
		currentPlatforms[p] = v
	}

	currentPlatforms[r.expectedPlatform] = r.possibleValue

	output := platformPrettyJson(currentPlatforms)
	var newPlatformIdx int
	for idx, out := range output {
		if strings.Contains(out, r.expectedPlatform) && strings.Contains(out, r.possibleValue) {
			newPlatformIdx = idx
		}
	}

	return printer.NewHighlightCodePrinter(output, []string{output[newPlatformIdx]}, nil).Print()
}
