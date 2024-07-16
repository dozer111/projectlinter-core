package composer_json_test

import (
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScriptsMerge(t *testing.T) {
	scripts := composer_json.NewScripts(
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
	)

	scripts2 := composer_json.NewScripts(
		map[string][]string{
			"post-install-cmd2": {"@auto-scripts"},
			"post-update-cmd2":  {"@auto-scripts"},
		},
		map[string]map[string]string{
			"auto-scripts2": {
				"cache:clear": "symfony-cmd",
			},
		},
		map[string]string{
			"php-cs-fixer2": "php-cs-fixer fix",
			"rector2":       "rector process",
		},
	)

	mergedScripts := scripts.Merge(scripts2)

	assert.Equal(t, 10, mergedScripts.Len())
}

func TestScriptsHas(t *testing.T) {
	scripts := composer_json.NewScripts(
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
	)

	t.Run("Array", func(t *testing.T) {
		hasValue, scriptType := scripts.Has("post-install-cmd")

		assert.True(t, hasValue)
		assert.Equal(t, scriptType, composer_json.TypeArray)
	})

	t.Run("Object", func(t *testing.T) {
		hasValue, scriptType := scripts.Has("auto-scripts")

		assert.True(t, hasValue)
		assert.Equal(t, scriptType, composer_json.TypeObject)
	})

	t.Run("String", func(t *testing.T) {
		hasValue, scriptType := scripts.Has("rector")

		assert.True(t, hasValue)
		assert.Equal(t, scriptType, composer_json.TypeString)
	})

	t.Run("value is absent", func(t *testing.T) {
		hasValue, scriptType := scripts.Has("this key is absent")

		assert.False(t, hasValue)
		assert.Equal(t, scriptType, composer_json.ScriptsType(0))
	})

}
