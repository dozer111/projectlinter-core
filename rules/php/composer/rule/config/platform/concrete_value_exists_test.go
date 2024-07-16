package platform_test

import (
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule/config/platform"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpecifiedPlatformExistsRulePassedWhilePlatformExists(t *testing.T) {
	r := platform.NewSpecifiedPlatformExistsRule(
		map[string]string{
			"php":            "asdasd",
			"symfony/bridge": "123456",
		},
		"php",
		"^5.4",
	)
	r.Validate()

	assert.True(t, r.IsPassed())
}

func TestSpecifiedPlatformExistsRuleFailedWhilePlatformIsAbsent(t *testing.T) {
	r := platform.NewSpecifiedPlatformExistsRule(
		map[string]string{
			"php":            "asdasd",
			"symfony/bridge": "123456",
		},
		"rector/rector",
		"^5.4",
	)
	r.Validate()

	assert.False(t, r.IsPassed())
}
