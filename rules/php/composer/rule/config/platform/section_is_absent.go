package platform

import (
	"encoding/json"
	"github.com/dozer111/projectlinter-core/printer"
	"github.com/dozer111/projectlinter-core/rules"
	"strings"
)

// ConfigPlatformIsAbsentRule - opposite rule to ConfigPlatformExistsRule
type ConfigPlatformIsAbsentRule struct {
	value    *map[string]string
	isPassed bool
}

var _ rules.Rule = (*ConfigPlatformIsAbsentRule)(nil)

func NewConfigPlatformIsAbsentRule(val *map[string]string) *ConfigPlatformIsAbsentRule {
	return &ConfigPlatformIsAbsentRule{
		value: val,
	}
}

func (r *ConfigPlatformIsAbsentRule) ID() string {
	return "composer.config.platform.absent"
}

func (r *ConfigPlatformIsAbsentRule) Title() string {
	return `"config.platform" is absent`
}

func (r *ConfigPlatformIsAbsentRule) Validate() {
	r.isPassed = r.value == nil
}

func (r *ConfigPlatformIsAbsentRule) IsPassed() bool {
	return r.isPassed
}

func (r *ConfigPlatformIsAbsentRule) FailedMessage() []string {
	type config struct {
		Platform map[string]string `json:"platform"`
	}

	currentPlatforms := make(map[string]string, len(*r.value))
	for p, v := range *r.value {
		currentPlatforms[p] = v
	}

	res := config{
		Platform: currentPlatforms,
	}

	b, _ := json.MarshalIndent(res, "", "	")
	prettyJson := strings.Split(string(b), "\n")
	prettyJson = prettyJson[1 : len(prettyJson)-1]

	return printer.NewCodeHintPrinter(
		nil,
		prettyJson,
		nil,
		nil,
	).Print()
}
