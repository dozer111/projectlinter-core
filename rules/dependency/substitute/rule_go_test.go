package substitute_test

import (
	"github.com/dozer111/projectlinter-core/rules/dependency/substitute"
	"github.com/dozer111/projectlinter-core/rules/golang/gomod/config"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

func TestSubstituteGOLibraryRulePass(t *testing.T) {
	t.Run("application/library does not has any library from configs", func(t *testing.T) {
		cfg := []substitute.Library{
			{
				Name: "github.com/agnivade/levenshtein",
			},
			{
				Name: "github.com/beorn7/perks",
			},
		}

		directDeps := config.NewGODependencies(0)
		directDeps.Add(
			config.NewGomodDependency(
				"github.com/chenzhuoyu/iasm",
				"v0.9.1",
				semver.MustParse("0.9.1"),
			),
		)

		indirectDeps := config.NewGODependencies(0)
		indirectDeps.Add(
			config.NewGomodDependency(
				"github.com/blendle/zapdriver",
				"v1.3.1",
				semver.MustParse("1.3.1"),
			),
		)

		substituteDependencies := goDependenciesToSubstituteDependencies(directDeps.Merge(indirectDeps))
		r := substitute.NewSubstituteGOLibraryRule("gomod", cfg, substituteDependencies)
		r.Validate()

		assert.True(t, r.IsPassed())
	})
}

func TestSubstituteGOLibraryRuleFail(t *testing.T) {
	t.Run("application require at least one library from configs", func(t *testing.T) {
		configs := []substitute.Library{
			{
				Name: "github.com/chenzhuoyu/iasm",
			},
		}

		directDeps := config.NewGODependencies(0)
		directDeps.Add(
			config.NewGomodDependency(
				"github.com/chenzhuoyu/iasm",
				"v0.9.1",
				semver.MustParse("0.9.1"),
			),
		)

		indirectDeps := config.NewGODependencies(0)
		substituteDependencies := goDependenciesToSubstituteDependencies(directDeps.Merge(indirectDeps))

		r := substitute.NewSubstituteGOLibraryRule("gomod", configs, substituteDependencies)
		r.Validate()

		assert.False(t, r.IsPassed())
	})
}

func TestSubstituteLibraryGORuleHandleVersionSuffix(t *testing.T) {
	// In GO, it is normal practice in libraries to specify the major version
	// example
	// github.com/Masterminds/sprig => v 1.0
	// github.com/Masterminds/sprig/v3 => v 3.0
	// We need to be sure that the versions in bump config and go.mod are comparable regardless of the presence/absence of the suffix

	t.Run("cfg:<name>/v2|gomod:<name>/<suffix>", func(t *testing.T) {
		configs := []substitute.Library{
			{
				Name: "github.com/Masterminds/sprig",
			},
			{
				Name: "github.com/pmezard/go-difflib",
			},
		}

		directDeps := config.NewGODependencies(0)
		directDeps.Add(config.NewGomodDependency(
			"github.com/Masterminds/sprig/v3",
			"v3.9.5",
			semver.MustParse("3.9.5"),
		),
		)

		indirectDeps := config.NewGODependencies(0)
		indirectDeps.Add(
			config.NewGomodDependency(
				"github.com/pmezard/go-difflib/v2",
				"v2.0.6",
				semver.MustParse("2.0.6"),
			),
		)

		substituteDependencies := goDependenciesToSubstituteDependencies(directDeps.Merge(indirectDeps))
		r := substitute.NewSubstituteGOLibraryRule("gomod", configs, substituteDependencies)
		r.Validate()

		assert.False(t, r.IsPassed())
	})

	t.Run("cfg:<name>/<suffix>/v2|gomod:<name>", func(t *testing.T) {
		configs := []substitute.Library{
			{
				Name: "github.com/Masterminds/sprig/v2",
			},
			{
				Name: "github.com/pmezard/go-difflib/v2",
			},
		}

		directDeps := config.NewGODependencies(0)
		directDeps.Add(config.NewGomodDependency(
			"github.com/Masterminds/sprig",
			"v3.9.5",
			semver.MustParse("3.9.5"),
		),
		)

		indirectDeps := config.NewGODependencies(0)
		indirectDeps.Add(
			config.NewGomodDependency(
				"github.com/pmezard/go-difflib",
				"v2.0.6",
				semver.MustParse("2.0.6"),
			),
		)

		substituteDependencies := goDependenciesToSubstituteDependencies(directDeps.Merge(indirectDeps))
		r := substitute.NewSubstituteGOLibraryRule("gomod", configs, substituteDependencies)
		r.Validate()

		assert.False(t, r.IsPassed())
	})

	t.Run(`cfg:<name>\/<suffix>/v2|gomod:<name>/<suffix>`, func(t *testing.T) {
		configs := []substitute.Library{
			{
				Name: "github.com/Masterminds/sprig/v2",
			},
			{
				Name: "github.com/pmezard/go-difflib/v2",
			},
		}

		directDeps := config.NewGODependencies(0)
		directDeps.Add(config.NewGomodDependency(
			"github.com/Masterminds/sprig/v3",
			"v3.9.5",
			semver.MustParse("3.9.5"),
		),
		)

		indirectDeps := config.NewGODependencies(0)
		indirectDeps.Add(
			config.NewGomodDependency(
				"github.com/pmezard/go-difflib/v2",
				"v2.0.6",
				semver.MustParse("2.0.6"),
			),
		)

		substituteDependencies := goDependenciesToSubstituteDependencies(directDeps.Merge(indirectDeps))
		r := substitute.NewSubstituteGOLibraryRule("gomod", configs, substituteDependencies)
		r.Validate()

		assert.False(t, r.IsPassed())
	})
}

func goDependenciesToSubstituteDependencies(dependencies *config.GomodDependencies) []substitute.Dependency {
	result := make([]substitute.Dependency, 0, dependencies.Count())

	for _, v := range dependencies.All() {
		for _, concreteDependency := range v {
			result = append(result, concreteDependency)
		}
	}

	return result
}
