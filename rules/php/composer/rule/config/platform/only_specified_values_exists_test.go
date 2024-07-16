package platform_test

import (
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule/config/platform"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnlySpecifiedPlatformExistsRulePassedWhilePlatformsAreSameToExpected(t *testing.T) {
	platformsInConfig := map[string]string{
		"php": "8.2",
	}

	expectedPlatfomrs := map[string]string{
		"php": "8.0",
	}

	r := platform.NewOnlySpecifiedPlatformExistsRule(platformsInConfig, expectedPlatfomrs)
	r.Validate()

	assert.True(t, r.IsPassed())
}

func TestOnlySpecifiedPlatformExistsRulePassedWhilePlatformsAreDifferFromExpected(t *testing.T) {
	platformsInConfig := map[string]string{
		"php": "8.2",
	}

	expectedPlatfomrs := map[string]string{
		"php":     "8.0",
		"ext-xml": "1.12",
	}

	r := platform.NewOnlySpecifiedPlatformExistsRule(platformsInConfig, expectedPlatfomrs)
	r.Validate()

	assert.False(t, r.IsPassed())
}
