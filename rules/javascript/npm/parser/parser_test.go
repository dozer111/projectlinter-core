package parser_test

import (
	"errors"
	"testing"

	utilTest "github.com/dozer111/projectlinter-core/util/test"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"github.com/dozer111/projectlinter-core/rules/javascript/npm/config/package_json"
	"github.com/dozer111/projectlinter-core/rules/javascript/npm/config/package_lock_json"
	javascriptNPMParser "github.com/dozer111/projectlinter-core/rules/javascript/npm/parser"
)

func TestParseSuccess(t *testing.T) {
	expectedPackageJSON := &package_json.RawPackageJSON{
		map[string]string{
			"express":       "^4.19.2",
			"lodash-master": "github:lodash/lodash#amd",
		},
		map[string]string{
			"eslint": "^9.9.1",
		},
	}

	expectedPackageLock := &package_lock_json.RawPackageLockJSON{
		map[string]*package_lock_json.RawNPMLockPackage{
			"": {"1.0.0"},
			"node_modules/@eslint-community/eslint-utils":                                  {"4.4.0"},
			"node_modules/@eslint-community/eslint-utils/node_modules/eslint-visitor-keys": {"3.4.3"},
			"node_modules/@eslint-community/regexpp":                                       {"4.11.0"},
			"node_modules/@eslint/config-array":                                            {"0.18.0"},
		},
	}

	p := javascriptNPMParser.NewParser(
		utilTest.PathInProjectLinter("testdata/success"),
	)
	actualPackageJSON, actualPackageLock, err := p.Parse()

	assert.Nil(t, err)
	if !cmp.Equal(expectedPackageJSON, actualPackageJSON) {
		assert.Fail(t, "package.json-s are not equal:\n"+cmp.Diff(expectedPackageJSON, actualPackageJSON))
	}
	if !cmp.Equal(expectedPackageLock, actualPackageLock) {
		assert.Fail(t, "package-lock.json-s are not equal:\n"+cmp.Diff(expectedPackageLock, actualPackageLock))
	}
}

func TestParserReturnErrorWhilePackageJsonIsAbsent(t *testing.T) {
	p := javascriptNPMParser.NewParser(
		utilTest.PathInProjectLinter("testdata/no_package_json"),
	)
	actualPackageJSON, actualPackageLock, err := p.Parse()

	assert.Nil(t, actualPackageJSON)
	assert.Nil(t, actualPackageLock)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, javascriptNPMParser.PackageJSONNotFound))
}

func TestParserReturnErrorWhilePackageLockIsAbsent(t *testing.T) {
	p := javascriptNPMParser.NewParser(
		utilTest.PathInProjectLinter("testdata/no_package_lock_json"),
	)
	actualPackageJson, actualPackageLock, err := p.Parse()

	assert.Nil(t, actualPackageJson)
	assert.Nil(t, actualPackageLock)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, javascriptNPMParser.PackageLockJSONNotFound))
}
