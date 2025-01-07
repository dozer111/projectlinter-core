package bump_test

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	utilTest "github.com/dozer111/projectlinter-core/util/test"

	"github.com/dozer111/projectlinter-core/rules/dependency/bump"

	"github.com/stretchr/testify/assert"
)

func TestParserSuccessCase(t *testing.T) {
	expectedConfigs := []bump.Library{
		{
			"my-private-git/doctrine-tool-lb",
			"1.3.0",
			nil,
			[]string{
				"dozer111",
			},
			nil,
		},
		{
			Name:    "my-private-git/logger",
			Version: "3.29.2",
			ResponsiblePersons: []string{
				"dondo",
			},
			Examples: []bump.Example{
				{
					ProjectName: "some-service",
					Programmer:  "dondo",
					Links: []string{
						"https://your_git.com/some-service/commits/497e11b2a6dda145a11e8384d48bc6fe97aa2de2",
					},
				},
				{
					"service2",
					"olof",
					[]string{
						"!IMPORTANT",
						"add lock.yaml",
						"configure in in cmd/main/init.go",
					},
					[]string{
						"https://your_git.com/some-service2/pull-requests/67/commits/0f2ce352e31c09aa79f13037c839b56ba701ba15",
					},
				},
			},
		},
		{
			"my-private-git/rector-rules",
			"2.0",
			nil,
			[]string{
				"dozer111",
			},
			[]bump.Example{
				{
					"some-service",
					"dozer111",
					nil,
					[]string{
						"https://your_git.com/service1/commits/228b2749cdec314558f04284b836b9d609f89e9f#rector.php",
					},
				},
			},
		},
	}

	parser := bump.NewParser(
		utilTest.PathInProjectLinter("testdata/success"),
	)
	actualConfigs, err := parser.Parse()

	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(expectedConfigs, actualConfigs))
}

func TestParserReturnErrorInPathToConfigsIsNotExists(t *testing.T) {
	p := bump.NewParser(
		utilTest.PathInProjectLinter("testdata/notExistingDirectory"),
	)
	actualConfigs, err := p.Parse()

	assert.Nil(t, actualConfigs)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, bump.PathToConfigsIsInvalid))

}

func TestParserReturnEmptySliceIfPathIsCorrectButDirectoryIsEmpty(t *testing.T) {
	p := bump.NewParser(
		utilTest.PathInProjectLinter("testdata/no_configs"),
	)
	actualConfigs, err := p.Parse()

	assert.Nil(t, err)
	assert.True(t, len(actualConfigs) == 0)
}

func TestParserReturnErrorWhileConfigDoesNotFitSchema(t *testing.T) {
	cases := []string{
		"examples_committee_is_absent",
		"examples_links_is_absent",
		"name_is_absent",
		"version_has_wrong_format",
		"examples_serviceName_is_absent",
		"version_is_absent",
	}

	for _, testCase := range cases {
		t.Run(strings.ReplaceAll(testCase, "_", " "), func(t *testing.T) {
			p := bump.NewParser(
				utilTest.PathInProjectLinter("testdata/config_does_not_fit_schema/" + testCase),
			)
			actualConfigs, err := p.Parse()

			assert.Nil(t, actualConfigs)
			assert.Error(t, err)
			assert.True(t, errors.Is(err, bump.ConfigDoesNotFitSchema))
		})
	}
}
