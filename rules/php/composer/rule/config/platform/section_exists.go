package platform

import (
	"fmt"
	"github.com/dozer111/projectlinter-core/printer"
	"github.com/dozer111/projectlinter-core/rules"
	"strings"
)

// ConfigPlatformExistsRule is one example when rule.SectionExistsRule is not suitable
//
// rule.SectionExistsRule is used when the section value is primitive.
// config is complex json object
type ConfigPlatformExistsRule struct {
	value         *map[string]string
	proposedValue map[string]string
	isPassed      bool
}

var _ rules.Rule = (*ConfigPlatformExistsRule)(nil)

func NewConfigPlatformExistsRule(
	val *map[string]string,
	proposedValue map[string]string,
) *ConfigPlatformExistsRule {
	return &ConfigPlatformExistsRule{
		value:         val,
		proposedValue: proposedValue,
	}
}

func (r *ConfigPlatformExistsRule) ID() string {
	return "composer.config.platform.exists"
}

func (r *ConfigPlatformExistsRule) Title() string {
	return `"config.platform" exists`
}

func (r *ConfigPlatformExistsRule) Validate() {
	r.isPassed = r.value != nil
}

func (r *ConfigPlatformExistsRule) IsPassed() bool {
	return r.isPassed
}

func (r *ConfigPlatformExistsRule) FailedMessage() []string {
	var output []string
	output = append(output, `"platform": {`)
	for platform, platformVersion := range r.proposedValue {
		output = append(output, fmt.Sprintf("\t\"%s\": \"%s\",", platform, platformVersion))
	}
	output = append(output, `},`)

	output[len(output)-2] = strings.TrimSuffix(output[len(output)-2], ",")

	return printer.NewCodeHintPrinter(
		nil,
		nil,
		output,
		nil,
	).Print()
}
