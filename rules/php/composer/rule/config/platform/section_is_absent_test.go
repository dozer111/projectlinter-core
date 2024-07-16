package platform_test

import (
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule/config/platform"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigPlatformAbsentRulePassedWhileSectionIsAbsent(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		r := platform.NewConfigPlatformIsAbsentRule(
			nil,
		)
		r.Validate()

		assert.True(t, r.IsPassed())
	})

	t.Run("non-initialized variable", func(t *testing.T) {
		var val *map[string]string
		r := platform.NewConfigPlatformIsAbsentRule(
			val,
		)
		r.Validate()

		assert.True(t, r.IsPassed())
	})
}

func TestConfigPlatformAbsentRuleFailedWhileSectionExists(t *testing.T) {
	r := platform.NewConfigPlatformIsAbsentRule(
		&map[string]string{},
	)
	r.Validate()

	assert.False(t, r.IsPassed())
}
