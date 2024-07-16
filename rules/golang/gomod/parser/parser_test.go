package parser_test

import (
	"errors"
	gomodParser "github.com/dozer111/projectlinter-core/rules/golang/gomod/parser"
	utilTest "github.com/dozer111/projectlinter-core/util/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	parser := gomodParser.NewParser(
		utilTest.PathInProjectLinter("testdata/success"),
	)

	mf, err := parser.Parse()

	assert.NoError(t, err)
	assert.NotNil(t, mf)
}

func TestParserReturnErrorIfGoModIsAbsent(t *testing.T) {
	parser := gomodParser.NewParser(
		utilTest.PathInProjectLinter("testdata/no_go_mod"),
	)

	mf, err := parser.Parse()

	assert.Nil(t, mf)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gomodParser.GoModIsAbsent))
}
