package scripts_test

import (
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule/scripts"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScriptsExistsRulePassedWhileSectionExists(t *testing.T) {
	r := scripts.NewScriptsExistsRule(
		&composer_json.Scripts{},
		composer_json.RawScripts{},
	)
	r.Validate()

	assert.True(t, r.IsPassed())
}

func TestScriptsExistsRuleFailedWhileSectionIsEmpty(t *testing.T) {
	r := scripts.NewScriptsExistsRule(
		nil,
		composer_json.RawScripts{},
	)
	r.Validate()

	assert.False(t, r.IsPassed())
}
