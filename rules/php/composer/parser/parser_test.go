package parser_test

import (
	"errors"
	"testing"

	utilTest "github.com/dozer111/projectlinter-core/util/test"
	"github.com/google/go-cmp/cmp"

	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_lock"
	"github.com/dozer111/projectlinter-core/rules/php/composer/parser"

	"github.com/stretchr/testify/assert"
)

func TestParseSuccess(t *testing.T) {
	expectedComposerJson := &composer_json.RawComposerJson{
		"private-git/some-project",
		pointer("project"),
		pointer("some php project"),
		pointer("proprietary"),
		nil,
		nil,
		map[string]string{
			"php":                                 "^8.2",
			"ext-openssl":                         "*",
			"doctrine/doctrine-migrations-bundle": "^3.2.2",
			"qandidate/symfony-json-request-transformer": "^2.2.0",
			"symfony/validator":                          "^6.2.10",
			"symfony/yaml":                               "^6.2.10",
		},
		map[string]string{
			"friendsofphp/php-cs-fixer": "^3.16.0",
			"symfony/browser-kit":       "^6.2.7",
			"symfony/phpunit-bridge":    "^6.2.10",
		},
		&composer_json.RawComposerJsonAutoloadSection{
			map[string]string{
				"PrivateGit\\SomeProject\\": "src/Service/",
				"Infrastructure\\":          "src/Infrastructure/",
			},
		},
		&composer_json.RawComposerJsonAutoloadSection{
			map[string]string{
				"PrivateGit\\SomeProject\\Tests\\": "tests/",
			},
		},
		&composer_json.RawComposerJsonConfigSection{
			pointer(true),
			pointer("dev"),
			pointer(map[string]string{"php": "8.2"}),
			pointer(map[string]bool{
				"symfony/flex":       true,
				"php-http/discovery": true,
			}),
		},
		map[string]string{
			"symfony/symfony": "*",
		},
		&composer_json.RawScripts{
			map[string][]string{
				"post-install-cmd": {"@auto-scripts"},
				"post-update-cmd":  {"@auto-scripts"},
			},
			map[string]map[string]string{
				"auto-scripts": {
					"cache:clear": "symfony-cmd",
				},
			},
			map[string]string{
				"php-cs-fixer": "php-cs-fixer fix",
				"rector":       "rector process",
			},
		},
	}
	expectedComposerLock := &composer_lock.RawComposerLock{
		map[string]string{
			"php":         "^8.2",
			"ext-openssl": "*",
		},
		map[string]string{
			"php": "8.2",
		},
		[]*composer_lock.RawComposerLockPackage{
			{
				"doctrine/cache",
				"2.2.0",
			},
			{
				"doctrine/dbal",
				"4.0.4",
			},
		},
		[]*composer_lock.RawComposerLockAlias{},
		[]*composer_lock.RawComposerLockPackage{
			{
				"clue/ndjson-react",
				"v1.3.0",
			},
			{
				"composer/pcre",
				"3.1.4",
			},
			{
				"composer/semver",
				"3.4.2",
			},
		},
	}

	p := parser.NewParser(
		utilTest.PathInProjectLinter("testdata/success"),
	)
	actualComposerJson, actualComposerLock, err := p.Parse()

	assert.Nil(t, err)
	if !cmp.Equal(expectedComposerJson, actualComposerJson) {
		assert.Fail(t, "composer.json-s are not equal:\n"+cmp.Diff(expectedComposerJson, actualComposerJson))
	}
	if !cmp.Equal(expectedComposerLock, actualComposerLock) {
		assert.Fail(t, "composer.lock-s are not equal:\n"+cmp.Diff(expectedComposerLock, actualComposerLock))
	}
}

func TestParserReturnErrorWhileComposerJsonIsAbsent(t *testing.T) {
	p := parser.NewParser(
		utilTest.PathInProjectLinter("testdata/no_composer_json"),
	)
	actualComposerJson, actualComposerLock, err := p.Parse()

	assert.Nil(t, actualComposerJson)
	assert.Nil(t, actualComposerLock)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, parser.ComposerJsonNotFound))
}

func TestParserReturnErrorWhileComposerLockIsAbsent(t *testing.T) {
	p := parser.NewParser(
		utilTest.PathInProjectLinter("testdata/no_composer_lock"),
	)
	actualComposerJson, actualComposerLock, err := p.Parse()

	assert.Nil(t, actualComposerJson)
	assert.Nil(t, actualComposerLock)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, parser.ComposerLockNotFound))
}

func pointer[T any](val T) *T {
	return &val
}
