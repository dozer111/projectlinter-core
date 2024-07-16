package rule_test

import (
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

func TestDependencyIsAbsentPassedWhileDependencyIsAbsent(t *testing.T) {
	depsWithSymfonyConsole := composer_json.NewComposerDependencies()
	depsWithSymfonyConsole.Add(composer_json.NewComposerDependency(
		"symfony/console",
		"^6.2.10",
		semver.MustParse("6.2.10"),
	),
	)

	r := rule.NewComposerDependencyIsAbsentRule(
		depsWithSymfonyConsole,
		"doctrine/doctrine-bundle",
		nil,
	)
	r.Validate()

	assert.True(t, r.IsPassed())
}

func TestDependencyIsAbsentFailWhileDependencyExists(t *testing.T) {
	depsWithSymfonyConsole := composer_json.NewComposerDependencies()
	depsWithSymfonyConsole.Add(
		composer_json.NewComposerDependency(
			"symfony/console",
			"^6.2.10",
			semver.MustParse("6.2.10"),
		),
	)

	r := rule.NewComposerDependencyIsAbsentRule(
		depsWithSymfonyConsole,
		"symfony/console",
		nil,
	)
	r.Validate()

	assert.False(t, r.IsPassed())
}
