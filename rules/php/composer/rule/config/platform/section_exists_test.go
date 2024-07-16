package platform_test

import (
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule/config/platform"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigPlatformExistsRulePassedWhileSectionExists(t *testing.T) {
	r := platform.NewConfigPlatformExistsRule(
		&map[string]string{},
		make(map[string]string),
	)
	r.Validate()

	assert.True(t, r.IsPassed())

}

func TestConfigPlatformExistsRuleFailedWhileSectionIsEmpty(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		r := platform.NewConfigPlatformExistsRule(
			nil,
			make(map[string]string),
		)
		r.Validate()

		assert.False(t, r.IsPassed())
	})

	t.Run("non-initialized variable", func(t *testing.T) {
		var val *map[string]string
		r := platform.NewConfigPlatformExistsRule(
			val,
			make(map[string]string),
		)
		r.Validate()

		assert.False(t, r.IsPassed())
	})
}
