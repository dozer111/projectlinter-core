package substitute_test

import (
	"github.com/dozer111/projectlinter-core/rules/dependency/substitute"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

func TestSubstitutePHPLibraryRulePass(t *testing.T) {
	t.Run("application/library does not has any library from configs", func(t *testing.T) {
		config := []substitute.Library{
			{
				Name: "mongo/mongo",
			},
			{
				Name: "symfony/flex",
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

		depsDev := composer_json.NewComposerDependencies()
		depsDev.Add(
			composer_json.NewComposerDependency(
				"rector/rector",
				"^0.18",
				semver.MustParse("0.18.2"),
			),
		)

		r := substitute.NewSubstitutePHPLibraryRule(
			"composer",
			config,
			phpDependenciesToSubstituteDependencies(deps.Merge(depsDev)),
		)
		r.Validate()

		assert.True(t, r.IsPassed())
	})
}

func TestSubstitutePHPLibraryRuleFail(t *testing.T) {
	t.Run("application require at least one library from configs", func(t *testing.T) {
		configs := []substitute.Library{
			{
				Name: "mongodb/mongo",
			},
		}

		deps := composer_json.NewComposerDependencies()
		deps.Add(
			composer_json.NewComposerDependency(
				"mongodb/mongo",
				"^1.15",
				semver.MustParse("1.15"),
			),
		)

		depsEmpty := composer_json.NewComposerDependencies()

		r := substitute.NewSubstitutePHPLibraryRule(
			"composer",
			configs,
			phpDependenciesToSubstituteDependencies(deps.Merge(depsEmpty)),
		)
		r.Validate()

		assert.False(t, r.IsPassed())
	})
}

func phpDependenciesToSubstituteDependencies(dependencies *composer_json.ComposerDependencies) []substitute.Dependency {
	result := make([]substitute.Dependency, 0, dependencies.Count())

	for _, v := range dependencies.All() {
		result = append(result, v)
	}

	return result
}
