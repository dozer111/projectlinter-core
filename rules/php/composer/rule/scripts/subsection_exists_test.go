package scripts_test

import (
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule/scripts"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScriptsSubsectionExistsRuleSuccess(t *testing.T) {
	t.Run("section is the same valueType", func(t *testing.T) {
		r := scripts.NewScriptsSubsectionExistsRule(
			"rector",
			composer_json.NewScripts(
				nil,
				nil,
				map[string]string{
					"php-cs-fixer": "php-cs-fixer fix",
					"rector":       "rector process",
				}),
			*composer_json.NewScripts(
				nil,
				nil,
				map[string]string{"rector": "rector process"},
			),
		)

		r.Validate()

		assert.True(t, r.IsPassed())
	})

	t.Run("section is different valueType", func(t *testing.T) {
		r := scripts.NewScriptsSubsectionExistsRule(
			"rector",
			composer_json.NewScripts(
				map[string][]string{
					"rector": {"rector", "process"},
				},
				nil,
				nil,
			),
			*composer_json.NewScripts(
				nil,
				nil,
				map[string]string{"rector": "rector process"},
			),
		)

		r.Validate()

		assert.True(t, r.IsPassed())
	})

}

func TestScriptsSubsectionExistsRuleFail(t *testing.T) {
	r := scripts.NewScriptsSubsectionExistsRule(
		"rector",
		composer_json.NewScripts(
			nil,
			nil,
			map[string]string{
				"php-cs-fixer2": "php-cs-fixer fix",
				"rector2":       "rector process",
			}),
		*composer_json.NewScripts(
			nil,
			nil,
			map[string]string{"rector": "rector process"},
		),
	)

	r.Validate()

	assert.False(t, r.IsPassed())
}
