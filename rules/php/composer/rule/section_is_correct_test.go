package rule_test

import (
	"strings"
	"testing"

	composerCustomType "github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json/type"
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

func TestSectionIsCorrect_HandleBoolString(t *testing.T) {
	// check that SectionHasCorrectValueRule can work with BoolString
	// and has correct failed message
	t.Run("string", func(t *testing.T) {
		t.Run("success(values are equal)", func(t *testing.T) {
			val := composerCustomType.BoolString{StrVal: "mirage"}
			r := rule.NewSectionHasCorrectValueRule("type", val, val)
			r.Validate()

			assert.True(t, r.IsPassed())
		})

		t.Run("fail(values are different)", func(t *testing.T) {
			t.Run("different string", func(t *testing.T) {
				val := composerCustomType.BoolString{StrVal: "mirage"}
				val2 := composerCustomType.BoolString{StrVal: "dust2"}
				r := rule.NewSectionHasCorrectValueRule("type", val, val2)
				r.Validate()

				assert.False(t, r.IsPassed())
				msg := strings.Join(r.FailedMessage(), "\n")
				assert.Equal(t, msg, "\u001B[31m\"type\": \"mirage\",\u001B[0m\n\u001B[32m\"type\": \"dust2\",\u001B[0m")
			})

			t.Run("value is bool", func(t *testing.T) {
				val := composerCustomType.BoolString{StrVal: "mirage"}
				val2 := composerCustomType.BoolString{IsBool: true, BoolVal: false}
				r := rule.NewSectionHasCorrectValueRule("type", val, val2)
				r.Validate()

				assert.False(t, r.IsPassed())
				msg := strings.Join(r.FailedMessage(), "\n")
				assert.Equal(t, msg, "\u001B[31m\"type\": \"mirage\",\u001B[0m\n\u001B[32m\"type\": false,\u001B[0m")
			})
		})
	})

	t.Run("bool", func(t *testing.T) {
		t.Run("success(values are equal)", func(t *testing.T) {
			val := composerCustomType.BoolString{IsBool: true, BoolVal: true}
			r := rule.NewSectionHasCorrectValueRule("type", val, val)
			r.Validate()

			assert.True(t, r.IsPassed())
		})

		t.Run("fail(values are different)", func(t *testing.T) {
			t.Run("different bool", func(t *testing.T) {
				val := composerCustomType.BoolString{IsBool: true, BoolVal: true}
				val2 := composerCustomType.BoolString{IsBool: true, BoolVal: false}
				r := rule.NewSectionHasCorrectValueRule("type", val, val2)
				r.Validate()

				assert.False(t, r.IsPassed())
				msg := strings.Join(r.FailedMessage(), "\n")
				assert.Equal(t, msg, "\u001B[31m\"type\": true,\u001B[0m\n\u001B[32m\"type\": false,\u001B[0m")
			})

			t.Run("value is string", func(t *testing.T) {
				val := composerCustomType.BoolString{IsBool: true, BoolVal: false}
				val2 := composerCustomType.BoolString{StrVal: "vertigo"}
				r := rule.NewSectionHasCorrectValueRule("type", val, val2)
				r.Validate()

				assert.False(t, r.IsPassed())
				msg := strings.Join(r.FailedMessage(), "\n")
				assert.Equal(t, msg, "\u001B[31m\"type\": false,\u001B[0m\n\u001B[32m\"type\": \"vertigo\",\u001B[0m")
			})
		})
	})
}
