package rule_test

import (
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSectionIsCorrectPassedWhenValueIsCorrect(t *testing.T) {
	const value = "project"
	r := rule.NewSectionHasCorrectValueRule("type", value, value)
	r.Validate()

	assert.True(t, r.IsPassed())
}

func TestSectionIsCorrectFailed(t *testing.T) {
	cases := []struct {
		description   string
		expectedValue string
		actualValue   string
	}{
		{
			"expected value != actual",
			"dondo",
			"xvost",
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			r := rule.NewSectionHasCorrectValueRule(
				"type",
				c.expectedValue,
				c.actualValue,
			)
			r.Validate()

			assert.False(t, r.IsPassed())
		})
	}
}
