package bump_test

import (
	"testing"

	"github.com/dozer111/projectlinter-core/rules/javascript/npm/config"

	"github.com/dozer111/projectlinter-core/rules/dependency/bump"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

func TestBumpJavascriptNPMLibraryRule(t *testing.T) {
	t.Run("all the libraries has higher version", func(t *testing.T) {
		configs := []bump.Library{
			{
				Name:    "express",
				Version: "4.0",
			},
			{
				Name:    "lodash",
				Version: "4.0",
			},
		}

		deps := config.NewNPMDependencies(1)

		deps.Add(config.NewNPMDependency(
			"node_modules/express",
			"^4.2",
			semver.MustParse("4.2.19"),
		),
		)

		depsDev := config.NewNPMDependencies(1)
		depsDev.Add(
			config.NewNPMDependency(
				"node_modules/lodash",
				"^4.1",
				semver.MustParse("4.3.15"),
			),
		)

		r := bump.NewBumpJavascriptNPMibraryRule(
			"javascript_npm",
			configs,
			javascriptNPMDependenciesToBumpDependencies(deps.Merge(depsDev)),
		)
		r.Validate()

		assert.True(t, r.IsPassed())
	})
}

func TestBumpJavascriptNPMLibraryRule_Fail(t *testing.T) {
	t.Run("at least one of dependency is lower than in config", func(t *testing.T) {
		configs := []bump.Library{
			{
				Name:    "express",
				Version: "4.0",
			},
		}

		deps := config.NewNPMDependencies(1)
		deps.Add(
			config.NewNPMDependency(
				"express",
				"^3",
				semver.MustParse("3.7.0"),
			),
		)

		depsEmpty := config.NewNPMDependencies(1)

		r := bump.NewBumpJavascriptNPMibraryRule(
			"javascript_npm",
			configs,
			javascriptNPMDependenciesToBumpDependencies(deps.Merge(depsEmpty)),
		)
		r.Validate()

		assert.False(t, r.IsPassed())
	})
}

func javascriptNPMDependenciesToBumpDependencies(dependencies *config.NPMDependencies) []bump.Dependency {
	result := make([]bump.Dependency, 0, dependencies.Count())

	for _, v := range dependencies.All() {
		result = append(result, v)
	}

	return result
}
