package bump_test

import (
	"github.com/dozer111/projectlinter-core/rules/dependency/bump"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

func TestBumpLibraryPHPRulePass(t *testing.T) {
	t.Run("all the libraries has higher version", func(t *testing.T) {
		configs := []bump.Library{
			{
				Name:    "symfony/console",
				Version: "6.0",
			},
			{
				Name:    "rector/rector",
				Version: "0.17",
			},
		}

		deps := composer_json.NewComposerDependencies()

		deps.Add(composer_json.NewComposerDependency(
			"symfony/console",
			"^6.6",
			semver.MustParse("6.6"),
		),
		)

		depsDev := composer_json.NewComposerDependencies()
		depsDev.Add(
			composer_json.NewComposerDependency(
				"rector/rector",
				"^0.18",
				semver.MustParse("0.18.2"),
			),
		)

		r := bump.NewBumpPHPLibraryRule("composer", configs, phpDependenciesToBumpDependencies(deps.Merge(depsDev)))
		r.Validate()

		assert.True(t, r.IsPassed())
	})
}

func TestBumpLibraryPHPRuleFail(t *testing.T) {
	t.Run("at least one of dependency is lower than in config", func(t *testing.T) {
		configs := []bump.Library{
			{
				Name:    "symfony/console",
				Version: "7.0",
			},
		}

		deps := composer_json.NewComposerDependencies()
		deps.Add(
			composer_json.NewComposerDependency(
				"symfony/console",
				"^6.6",
				semver.MustParse("6.6"),
			),
		)

		depsEmpty := composer_json.NewComposerDependencies()

		r := bump.NewBumpPHPLibraryRule("composer", configs, phpDependenciesToBumpDependencies(deps.Merge(depsEmpty)))
		r.Validate()

		assert.False(t, r.IsPassed())
	})
}

func phpDependenciesToBumpDependencies(dependencies *composer_json.ComposerDependencies) []bump.Dependency {
	result := make([]bump.Dependency, 0, dependencies.Count())

	for _, v := range dependencies.All() {
		result = append(result, v)
	}

	return result
}
