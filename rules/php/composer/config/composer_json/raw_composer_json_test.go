package composer_json_test

import (
	"encoding/json"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	utilJSON "github.com/dozer111/projectlinter-core/util/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRawScriptsUnmarshal(t *testing.T) {
	scriptsSection := `
{
  "post-install-cmd": [
    "@auto-scripts"
  ],
  "post-update-cmd": [
    "@auto-scripts"
  ],
  "auto-scripts": {
    "cache:clear": "symfony-cmd"
  },
  "php-cs-fixer": "php-cs-fixer fix",
  "rector": "rector process"
}
`

	var actualValue composer_json.RawScripts
	err := json.Unmarshal([]byte(scriptsSection), &actualValue)

	require.NoError(t, err)

	expectedValue := composer_json.RawScripts{
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
	}

	require.Equal(t, expectedValue, actualValue)
}

func TestRawScriptsMarshal(t *testing.T) {
	scripts := composer_json.RawScripts{
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
	}

	actualJSON, err := json.Marshal(&scripts)
	require.NoError(t, err)

	expectedJSON := `
{
  "post-install-cmd": [
    "@auto-scripts"
  ],
  "post-update-cmd": [
    "@auto-scripts"
  ],
  "auto-scripts": {
    "cache:clear": "symfony-cmd"
  },
  "php-cs-fixer": "php-cs-fixer fix",
  "rector": "rector process"
}
`

	require.True(t, utilJSON.JSONsAreEqual([]byte(expectedJSON), actualJSON))
}
