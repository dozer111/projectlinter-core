package scripts

import (
	"encoding/json"
	"github.com/dozer111/projectlinter-core/printer"
	"github.com/dozer111/projectlinter-core/rules"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"strings"
)

// ScriptsExistsRule is one example when rule.SectionExistsRule is not suitable
//
// rule.SectionExistsRule is used when the section value is primitive.
// scripts is a complex structure that can simultaneously contain different types of data
type ScriptsExistsRule struct {
	value         *composer_json.Scripts
	proposedValue composer_json.RawScripts
	isPassed      bool
}

var _ rules.Rule = (*ScriptsExistsRule)(nil)

func NewScriptsExistsRule(
	val *composer_json.Scripts,
	proposedValue composer_json.RawScripts,
) *ScriptsExistsRule {
	return &ScriptsExistsRule{
		value:         val,
		proposedValue: proposedValue,
	}
}

func (r *ScriptsExistsRule) ID() string {
	return "composer.scripts.exists"
}

func (r *ScriptsExistsRule) Title() string {
	return `"config.scripts" exists`
}

func (r *ScriptsExistsRule) Validate() {
	r.isPassed = r.value != nil
}

func (r *ScriptsExistsRule) IsPassed() bool {
	return r.isPassed
}

func (r *ScriptsExistsRule) FailedMessage() []string {
	type section struct {
		Scripts *composer_json.RawScripts `json:"scripts"`
	}

	sectionJSON, err := json.MarshalIndent(section{&r.proposedValue}, "", "  ")
	if err != nil {
		panic("cannot marshal composer_json.Scripts")
	}

	jsonAsStringSlice := strings.Split(string(sectionJSON), "\n")
	jsonAsStringSlice = jsonAsStringSlice[1 : len(jsonAsStringSlice)-1]

	return printer.NewCodeHintPrinter(
		[]string{"..."},
		nil,
		jsonAsStringSlice,
		[]string{"..."},
	).Print()
}
