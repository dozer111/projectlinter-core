package substitute_test

import (
	"errors"
	"github.com/dozer111/projectlinter-core/rules/dependency/substitute"
	utilTest "github.com/dozer111/projectlinter-core/util/test"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserSuccessCase(t *testing.T) {
	expectedConfigs := []substitute.Library{
		{
			"code-tool/doctrine-jaeger-symfony-bridge",
			"code-tool/doctrine-dbal-jaeger",
			nil,
			nil,
			[]substitute.Example{
				{
					"auth-sv",
					"dozer111",
					[]string{
						"https://your_git.com/auth-sv/commits/c836698e64149412fab971e2252396e9370394f7",
						"https://your_git.com/auth-sv/commits/0ce7f3c9ee28c28d7ec96df459e45b443f3c16b6",
					},
				},
			},
		},
		{
			"private-workspace/log-lb",
			"private-workspace/new-log-lib-lb",
			nil,
			nil,
			nil,
		},
	}

	p := substitute.NewParser(
		utilTest.PathInProjectLinter("testdata/success"),
	)
	actualConfigs, err := p.Parse()

	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(expectedConfigs, actualConfigs))
}

func TestParserReturnErrorInPathToConfigsIsNotExists(t *testing.T) {

	p := substitute.NewParser(
		utilTest.PathInProjectLinter("testdata/notExistingDirectory"),
	)
	actualConfigs, err := p.Parse()

	assert.Nil(t, actualConfigs)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, substitute.PathToBumpConfigsIsInvalid))
}

func TestParserReturnEmptySliceIfPathIsCorrectButDirectoryIsEmpty(t *testing.T) {

	p := substitute.NewParser(
		utilTest.PathInProjectLinter("testdata/no_configs"),
	)
	actualConfigs, err := p.Parse()

	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual([]substitute.Library{}, actualConfigs))
}

func TestParserReturnErrorWhileConfigDoesNotFitSchema(t *testing.T) {
	cases := []string{
		"name_is_absent",
		"changeTo_is_absent",
		"example_comittee_is_absent",
		"example_links_is_absent",
		"example_serviceName_is_absent",
	}

	for _, testCase := range cases {
		t.Run(strings.ReplaceAll(testCase, "_", " "), func(t *testing.T) {
			p := substitute.NewParser(
				utilTest.PathInProjectLinter("testdata/config_does_not_fit_schema/" + testCase),
			)
			actualConfigs, err := p.Parse()

			assert.Nil(t, actualConfigs)
			assert.Error(t, err)
			assert.True(t, errors.Is(err, substitute.BumpConfigDoesNotFitSchema))
		})
	}
}
