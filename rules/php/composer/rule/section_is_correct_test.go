package rule_test

import (
	"testing"

	"github.com/dozer111/projectlinter-core/rules/php/composer/rule"

	"github.com/stretchr/testify/assert"
)

func TestSectionIsCorrect(t *testing.T) {
	t.Run("success cases", func(t *testing.T) {
		t.Run("#1 the single expected value equals actual", func(t *testing.T) {
			r := rule.NewSectionHasCorrectValueRule("type", "project", "project")
			r.Validate()

			assert.True(t, r.IsPassed())
		})

		t.Run("#2 one of expected values equals actual", func(t *testing.T) {
			r := rule.NewSectionHasCorrectValueRule(
				"type",
				"symfony-bundle",

				"library",
				"symfony-bundle",
			)
			r.Validate()

			assert.True(t, r.IsPassed())
		})
	})

	t.Run("failure cases", func(t *testing.T) {
		t.Run("#1 single expected value != actual", func(t *testing.T) {
			r := rule.NewSectionHasCorrectValueRule("type", "library", "project")
			r.Validate()

			assert.False(t, r.IsPassed())
		})

		t.Run("#2 actual value != any of expected", func(t *testing.T) {
			r := rule.NewSectionHasCorrectValueRule(
				"type",
				"library",

				"project",
				"symfony-bundle",
				"value",
			)
			r.Validate()

			assert.False(t, r.IsPassed())
		})

		t.Run("#3 expected value does not set", func(t *testing.T) {
			r := rule.NewSectionHasCorrectValueRule("type", "library")
			r.Validate()

			assert.False(t, r.IsPassed())
		})

	})
}
