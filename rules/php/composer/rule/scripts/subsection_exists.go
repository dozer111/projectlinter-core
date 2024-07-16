package scripts

import (
	"encoding/json"
	"fmt"
	"github.com/dozer111/projectlinter-core/printer"
	"github.com/dozer111/projectlinter-core/rules"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"strings"
)

// ScriptsSubsectionExistsRule
// check if there is a certain subsection in the scripts section
//
// Example: composer.json:
//
//	 "scripts": {
//	    "post-install-cmd": [
//	        "@auto-scripts"
//	    ],
//	    "post-update-cmd": [
//	        "@auto-scripts"
//	    ],
//	    "auto-scripts": {
//	        "cache:clear": "symfony-cmd"
//	    },
//	    "php-cs-fixer": "php-cs-fixer fix",
//	    "rector": "rector process"
//	}
//
// the rule will check if this section has a subsection "rector"
type ScriptsSubsectionExistsRule struct {
	subsectionName string
	value          *composer_json.Scripts
	proposedValue  composer_json.Scripts

	isPassed bool
}

var _ rules.Rule = (*ScriptsSubsectionExistsRule)(nil)

func NewScriptsSubsectionExistsRule(
	subsectionName string,
	val *composer_json.Scripts,
	proposedValue composer_json.Scripts,
) *ScriptsSubsectionExistsRule {
	return &ScriptsSubsectionExistsRule{
		subsectionName: subsectionName,
		value:          val,
		proposedValue:  proposedValue,
	}
}

func (r *ScriptsSubsectionExistsRule) ID() string {
	return fmt.Sprintf("composer.scripts.subsection.%s.exists", r.subsectionName)
}

func (r *ScriptsSubsectionExistsRule) Title() string {
	return fmt.Sprintf(`config.scripts subsection "%s" exists`, r.subsectionName)
}

func (r *ScriptsSubsectionExistsRule) Validate() {
	r.isPassed, _ = r.value.Has(r.subsectionName)
}

func (r *ScriptsSubsectionExistsRule) IsPassed() bool {
	return r.isPassed
}

func (r *ScriptsSubsectionExistsRule) FailedMessage() []string {
	jsonSectionWithProposedValue, err := json.MarshalIndent(composer_json.NewRawScriptsFromScripts(&r.proposedValue), "", "  ")
	if err != nil {
		panic("cannot marshal jsonSectionWithProposedValue")
	}

	jsonAsStringSlice := strings.Split(string(jsonSectionWithProposedValue), "\n")

	return printer.NewCodeHintPrinter(
		[]string{
			`"scripts": {`,
			`  ...`,
		},
		nil,
		jsonAsStringSlice[1:len(jsonAsStringSlice)-1],
		[]string{
			"}",
		},
	).
		Print()
}
