package dependency_test

import (
	"testing"

	"github.com/dozer111/projectlinter-core/util/path_provider"
	utilTest "github.com/dozer111/projectlinter-core/util/test"

	"github.com/dozer111/projectlinter-core/rules/dependency"

	"github.com/stretchr/testify/assert"
)

func TestDependencyPHPSet(t *testing.T) {
	pathProvider := path_provider.NewPathProvider(utilTest.PathInProjectLinter("testdata/dependency"))

	s := dependency.NewPHPDependencySet(pathProvider, nil, nil)
	errs := s.Init()
	assert.Equal(t, 0, len(errs))

	rules := s.Run().Resolve([]string{})

	assert.Equal(t, 2, len(rules))
	failRules := utilTest.AllSetRulesArePassed(rules)
	assert.True(t, len(failRules) == 0, "Some rules are failed:\n%v", failRules)
}
