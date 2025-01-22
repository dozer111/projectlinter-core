package rule_test

import (
	"regexp"
	"testing"

	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

func TestDependencyConstraintsAreValid(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		depsWithSymfonyConsole := composer_json.NewComposerDependencies()
		depsWithSymfonyConsole.Add(
			composer_json.NewComposerDependency(
				"symfony/console",
				"^8.1",
				semver.MustParse("8.1"),
			),
		)

		r := rule.NewDependenciesConstrainsAreValidRule(
			depsWithSymfonyConsole,
			func(dependency composer_json.ComposerDependency) bool {
				r1 := regexp.MustCompile(`^\^\d+$`)           // ^3
				r2 := regexp.MustCompile(`^\^\d+\.\d+$`)      // ^3.2
				r3 := regexp.MustCompile(`^\^\d+\.\d+\.\d+$`) // ^3.2.15

				return r1.MatchString(dependency.Constraint()) ||
					r2.MatchString(dependency.Constraint()) ||
					r3.MatchString(dependency.Constraint())
			},
			[]string{
				`Dependency constraint(except "ext-") must match one of patterns: ^d+.d+(^8.1) or ^d+.d+.d+(^8.1.2)`,
				"- ^d+$         ^8",
				"- ^d+.d+$      ^8.1",
				"- ^d+.d+.d+$   ^8.1.2",
			},
		)
		r.Validate()

		assert.True(t, r.IsPassed())
	})

	t.Run("fail", func(t *testing.T) {
		depsWithSymfonyConsole := composer_json.NewComposerDependencies()
		depsWithSymfonyConsole.Add(
			composer_json.NewComposerDependency(
				"ext-mongo",
				"*",
				nil,
			),
		)

		r := rule.NewDependenciesConstrainsAreValidRule(
			depsWithSymfonyConsole,
			func(dependency composer_json.ComposerDependency) bool {
				r1 := regexp.MustCompile(`^\^\d+$`)           // ^3
				r2 := regexp.MustCompile(`^\^\d+\.\d+$`)      // ^3.2
				r3 := regexp.MustCompile(`^\^\d+\.\d+\.\d+$`) // ^3.2.15

				return r1.MatchString(dependency.Constraint()) ||
					r2.MatchString(dependency.Constraint()) ||
					r3.MatchString(dependency.Constraint())
			},
			[]string{
				`Dependency constraint(except "ext-") must match one of patterns: ^d+.d+(^8.1) or ^d+.d+.d+(^8.1.2)`,
				"- ^d+$         ^8",
				"- ^d+.d+$      ^8.1",
				"- ^d+.d+.d+$   ^8.1.2",
			},
		)
		r.Validate()

		assert.False(t, r.IsPassed())
	})
}
