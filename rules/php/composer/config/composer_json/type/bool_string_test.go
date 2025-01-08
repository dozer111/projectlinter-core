package composerCustomType_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	composerCustomType "github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json/type"
)

func TestBoolString(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		var actualObj composerCustomType.BoolString
		err := json.Unmarshal([]byte(`"dondo"`), &actualObj)

		assert.NoError(t, err)
		assert.False(t, actualObj.IsBool)
		assert.Equal(t, actualObj.StrVal, "dondo")
	})

	t.Run("bool", func(t *testing.T) {
		var actualObj composerCustomType.BoolString
		err := json.Unmarshal([]byte(`false`), &actualObj)

		assert.NoError(t, err)
		assert.True(t, actualObj.IsBool)
		assert.Equal(t, actualObj.BoolVal, false)
	})
}
