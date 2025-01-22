package rule_test

import (
	"testing"

	composerCustomType "github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json/type"
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule"

	"github.com/stretchr/testify/assert"
)

func TestSectionIsAbsentPassedWhenSectionIsAbsent(t *testing.T) {
	t.Run("should pass on nil", func(t *testing.T) {
		var value *string
		r := rule.NewSectionIsAbsentRule("type", value)
		r.Validate()

		assert.True(t, r.IsPassed())
	})

	t.Run("should fail on", func(t *testing.T) {
		cases := map[string]any{
			// yes, it is zero value, but its not nil
			"empty string":     "",
			"not-empty string": "vailkamu",
			"boolstring":       composerCustomType.BoolString{StrVal: "mirage"},
		}

		for description, value := range cases {
			t.Run(description, func(t *testing.T) {
				r := rule.NewSectionIsAbsentRule("type", &value)
				r.Validate()

				assert.False(t, r.IsPassed())
			})
		}
	})

	t.Run("should handle type on fail", func(t *testing.T) {
		t.Run("string", func(t *testing.T) {
			strVal := ""
			r := rule.NewSectionIsAbsentRule("type", &strVal)
			r.Validate()
			assert.NotPanics(t, func() { r.FailedMessage() })
		})

		t.Run("bool", func(t *testing.T) {
			boolVal := false
			r := rule.NewSectionIsAbsentRule("type", &boolVal)
			r.Validate()
			assert.NotPanics(t, func() { r.FailedMessage() })
		})

		t.Run("boolString bool", func(t *testing.T) {
			boolStringBoolVal := composerCustomType.BoolString{BoolVal: false, IsBool: true}
			r := rule.NewSectionIsAbsentRule("type", &boolStringBoolVal)
			r.Validate()
			assert.NotPanics(t, func() { r.FailedMessage() })
		})

		t.Run("boolString string", func(t *testing.T) {
			boolStringStringVal := composerCustomType.BoolString{BoolVal: false, IsBool: true}
			r := rule.NewSectionIsAbsentRule("type", &boolStringStringVal)
			r.Validate()
			assert.NotPanics(t, func() { r.FailedMessage() })
		})
	})
}
