package rule_test

import (
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpecialDependencyExistsRulePassWhileDependencyExists(t *testing.T) {
	depsWithExtJson := composer_json.NewComposerDependencies()
	depsWithExtJson.Add(composer_json.NewComposerDependency("ext-json", "*", nil))

	r := rule.NewSpecialDependencyExistsRule(
		depsWithExtJson,
		"ext-json",
		"*",
		true,
	)

	r.Validate()

	assert.True(t, r.IsPassed())
}

func TestSpecialDependencyExistsRuleFailWhileDependencyIsAbsent(t *testing.T) {
	depsWithExtJson := composer_json.NewComposerDependencies()
	depsWithExtJson.Add(composer_json.NewComposerDependency("ext-json", "*", nil))

	r := rule.NewSpecialDependencyExistsRule(
		depsWithExtJson,
		"ext-xml",
		"*",
		true,
	)

	r.Validate()

	assert.False(t, r.IsPassed())
}
