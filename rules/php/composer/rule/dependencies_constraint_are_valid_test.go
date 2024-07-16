package rule_test

import (
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

func TestDependencyConstraintsAreValidSuccessCase(t *testing.T) {
	cases := []struct {
		desctiption string
		value       string
	}{
		{
			`d+`,
			"8",
		},
		{
			`d+.d+`,
			"8.1",
		},
		{
			`d+.d+.d+`,
			"8.1.15",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.desctiption, func(t *testing.T) {
			depsWithSymfonyConsole := composer_json.NewComposerDependencies()
			depsWithSymfonyConsole.Add(composer_json.NewComposerDependency(
				"symfony/console",
				"^"+testCase.value,
				semver.MustParse(testCase.value),
			),
			)
			r := rule.NewDependenciesConstrainsAreValidRule(depsWithSymfonyConsole, []string{})
			r.Validate()

			assert.True(t, r.IsPassed())
		})
	}
}

func TestDependencyConstraintsAreValidFailCase(t *testing.T) {
	depsWithWrongConstraint := composer_json.NewComposerDependencies()
	depsWithWrongConstraint.Add(
		composer_json.NewComposerDependency(
			"symfony/console",
			"dev-master",
			nil,
		),
	)

	r := rule.NewDependenciesConstrainsAreValidRule(depsWithWrongConstraint, []string{})
	r.Validate()

	assert.False(t, r.IsPassed())
}

func TestDependencyConstraintsAreValidIgnoreLibrariesInExceptionBlock(t *testing.T) {
	depsWithWrongConstraint := composer_json.NewComposerDependencies()
	depsWithWrongConstraint.Add(
		composer_json.NewComposerDependency(
			"symfony/console",
			"dev-master",
			nil,
		),
		composer_json.NewComposerDependency(
			"ext-json",
			"*",
			nil,
		),
	)

	r := rule.NewDependenciesConstrainsAreValidRule(depsWithWrongConstraint, []string{"symfony/console", "ext-json"})
	r.Validate()

	assert.True(t, r.IsPassed())
}
