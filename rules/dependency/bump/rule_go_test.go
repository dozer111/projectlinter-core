package bump_test

import (
	"testing"

	"github.com/dozer111/projectlinter-core/rules/golang/gomod/config"

	"github.com/dozer111/projectlinter-core/rules/dependency/bump"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

func TestBumpLibraryGORulePass(t *testing.T) {
	t.Run("all the libraries has higher version", func(t *testing.T) {
		configs := []bump.Library{
			{
				Name:    "github.com/pelletier/go-toml",
				Version: "1.5",
			},
			{
				Name:    "github.com/beorn7/perks",
				Version: "1.0",
			},
		}

		directDeps := config.NewGODependencies(0)
		directDeps.Add(
			config.NewGomodDependency(
				"github.com/pelletier/go-toml",
				"v1.9.5",
				semver.MustParse("1.9.5"),
				false,
			),
			config.NewGomodDependency(
				"github.com/beorn7/perks",
				"v1.0.1",
				semver.MustParse("1.0.1"),
				false,
			),
		)

		bumpDependencies := goDependenciesToBumpDependencies(directDeps)

		r := bump.NewBumpGOLibraryRule("gomod", configs, bumpDependencies)
		r.Validate()

		assert.True(t, r.IsPassed())
	})
}

func TestBumpLibraryGORule_Ignore_Indirect_Dependencies(t *testing.T) {
	// It is crucial to understand
	// Golang dependency rules MUST ONLY CHECK THE DIRECT DEPENDENCIES!
	// Why - because indirect dependencies are also exists in go.mod(instead as in other programming languages)
	// so, if the rule trigger on it - this is false positive check

	configs := []bump.Library{
		{
			Name:    "github.com/pelletier/go-toml",
			Version: "1.5",
		},
		{
			Name:    "github.com/beorn7/perks",
			Version: "2.0",
		},
	}

	directDeps := config.NewGODependencies(0)
	directDeps.Add(
		config.NewGomodDependency(
			"github.com/pelletier/go-toml",
			"v1.9.5",
			semver.MustParse("1.9.5"),
			false,
		),
	)

	indirectDeps := config.NewGODependencies(0)
	indirectDeps.Add(
		config.NewGomodDependency(
			"github.com/beorn7/perks",
			"v1.1.15",
			semver.MustParse("1.1.15"),
			true,
		),
	)

	bumpDependencies := goDependenciesToBumpDependencies(directDeps.Merge(indirectDeps))

	r := bump.NewBumpGOLibraryRule("gomod", configs, bumpDependencies)
	r.Validate()

	assert.True(t, r.IsPassed())
}

func TestBumpLibraryGORuleFail(t *testing.T) {
	t.Run("at least one of dependency is lower than in config", func(t *testing.T) {
		configs := []bump.Library{
			{
				Name:    "github.com/pkg/errors/v2",
				Version: "2.0",
			},
		}

		directDeps := config.NewGODependencies(0)
		directDeps.Add(
			config.NewGomodDependency(
				"github.com/pkg/errors",
				"v1.9.1",
				semver.MustParse("1.9"),
				false,
			),
		)

		bumpDependencies := goDependenciesToBumpDependencies(directDeps)
		r := bump.NewBumpGOLibraryRule("gomod", configs, bumpDependencies)
		r.Validate()

		assert.False(t, r.IsPassed())
	})

	t.Run("project has 2 versions of library and at least one of dependency is lower than in config", func(t *testing.T) {
		configs := []bump.Library{
			{
				Name:    "github.com/pkg/errors/v2",
				Version: "2.0",
			},
		}

		dependencies := config.NewGODependencies(2)
		dependencies.Add(
			config.NewGomodDependency(
				"github.com/pkg/errors/v3",
				"v3.2.0",
				semver.MustParse("3.2"),
				false,
			),
			config.NewGomodDependency(
				"github.com/pkg/errors/v3",
				"v2.17.0",
				semver.MustParse("2.17"),
				false,
			),
			config.NewGomodDependency(
				"github.com/pkg/errors",
				"v1.9.1",
				semver.MustParse("1.9"),
				false,
			),
		)

		bumpDependencies := goDependenciesToBumpDependencies(dependencies)
		r := bump.NewBumpGOLibraryRule("gomod", configs, bumpDependencies)
		r.Validate()

		assert.False(t, r.IsPassed())
	})
}

func TestBumpLibraryGORuleHandleVersionSuffix(t *testing.T) {
	// In GO, it is normal practice in libraries to specify the major version
	// example
	// github.com/Masterminds/sprig => v 1.0
	// github.com/Masterminds/sprig/v3 => v 3.0
	// We need to be sure that the versions in bump config and go.mod are comparable regardless of the presence/absence of the suffix

	t.Run("cfg:<name>/v2|gomod:<name>/<suffix>", func(t *testing.T) {
		configs := []bump.Library{
			{
				Name:    "github.com/Masterminds/sprig",
				Version: "2.5",
			},
			{
				Name:    "github.com/pmezard/go-difflib",
				Version: "1.0",
			},
		}

		directDeps := config.NewGODependencies(0)
		directDeps.Add(config.NewGomodDependency(
			"github.com/Masterminds/sprig/v3",
			"v3.9.5",
			semver.MustParse("3.9.5"),
			false,
		),
			config.NewGomodDependency(
				"github.com/pmezard/go-difflib/v2",
				"v2.0.6",
				semver.MustParse("2.0.6"),
				false,
			),
		)

		bumpDependencies := goDependenciesToBumpDependencies(directDeps)

		r := bump.NewBumpGOLibraryRule("gomod", configs, bumpDependencies)
		r.Validate()

		assert.True(t, r.IsPassed())
	})

	t.Run("cfg:<name>/<suffix>/v2|gomod:<name>", func(t *testing.T) {
		configs := []bump.Library{
			{
				Name:    "github.com/Masterminds/sprig/v2",
				Version: "2.5",
			},
			{
				Name:    "github.com/pmezard/go-difflib/v2",
				Version: "2.0",
			},
		}

		directDeps := config.NewGODependencies(0)
		directDeps.Add(
			config.NewGomodDependency(
				"github.com/Masterminds/sprig",
				"v3.9.5",
				semver.MustParse("3.9.5"),
				false,
			),
			config.NewGomodDependency(
				"github.com/pmezard/go-difflib",
				"v2.0.6",
				semver.MustParse("2.0.6"),
				false,
			),
		)

		bumpDependencies := goDependenciesToBumpDependencies(directDeps)

		r := bump.NewBumpGOLibraryRule("gomod", configs, bumpDependencies)
		r.Validate()

		assert.True(t, r.IsPassed())
	})

	t.Run(`cfg:<name>\/<suffix>/v2|gomod:<name>/<suffix>`, func(t *testing.T) {
		configs := []bump.Library{
			{
				Name:    "github.com/Masterminds/sprig/v2",
				Version: "2.5",
			},
			{
				Name:    "github.com/pmezard/go-difflib/v2",
				Version: "2.0",
			},
		}

		directDeps := config.NewGODependencies(0)
		directDeps.Add(config.NewGomodDependency(
			"github.com/Masterminds/sprig/v3",
			"v3.9.5",
			semver.MustParse("3.9.5"),
			false,
		),
			config.NewGomodDependency(
				"github.com/pmezard/go-difflib/v2",
				"v2.0.6",
				semver.MustParse("2.0.6"),
				false,
			),
		)

		bumpDependencies := goDependenciesToBumpDependencies(directDeps)

		r := bump.NewBumpGOLibraryRule("gomod", configs, bumpDependencies)
		r.Validate()

		assert.True(t, r.IsPassed())
	})
}

func goDependenciesToBumpDependencies(dependencies *config.GomodDependencies) []bump.Dependency {
	result := make([]bump.Dependency, 0, dependencies.Count())

	for _, v := range dependencies.All() {
		for _, concreteDependency := range v {
			result = append(result, concreteDependency)
		}
	}

	return result
}
