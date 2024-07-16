package rule_test

import (
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

func TestDependencyExistsPassedWhileDependencyExists(t *testing.T) {
	depsWithSymfonyConsole := composer_json.NewComposerDependencies()
	depsWithSymfonyConsole.Add(
		composer_json.NewComposerDependency(
			"symfony/console",
			"^6.2.10",
			semver.MustParse("6.2.10"),
		),
	)

	r := rule.NewComposerDependencyExistsRule(
		depsWithSymfonyConsole,
		"symfony/console",
		true,
		nil,
	)
	r.Validate()

	assert.True(t, r.IsPassed())
}

func TestDependencyExistsFailWhileDependencyIsAbsent(t *testing.T) {
	depsWithSymfonyConsole := composer_json.NewComposerDependencies()
	depsWithSymfonyConsole.Add(
		composer_json.NewComposerDependency(
			"symfony/console",
			"^6.2.10",
			semver.MustParse("6.2.10"),
		),
	)

	r := rule.NewComposerDependencyExistsRule(
		depsWithSymfonyConsole,
		"doctrine/doctrine-bundle",
		true,
		nil,
	)
	r.Validate()

	assert.False(t, r.IsPassed())
}
