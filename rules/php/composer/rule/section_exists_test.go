package rule_test

import (
	"testing"

	composerCustomType "github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json/type"
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule"

	"github.com/stretchr/testify/assert"
)

func TestSectionExists(t *testing.T) {
	t.Run("should pass on non-nil", func(t *testing.T) {
		testCases := map[string]any{
			"string":     "",
			"bool":       false,
			"boolstring": composerCustomType.BoolString{StrVal: "omange"},
		}

		for description, value := range testCases {
			t.Run(description, func(t *testing.T) {
				r := rule.NewSectionExistsRule("type", &value, value)
				r.Validate()

				assert.True(t, r.IsPassed())
			})
		}
	})

	t.Run("should fail on nil", func(t *testing.T) {
		r := rule.NewSectionExistsRule("type", nil, "project")
		r.Validate()

		assert.False(t, r.IsPassed())
	})

	t.Run("should handle types on failure", func(t *testing.T) {
		// when rule is failed - it must print correct message
		// this test case check that specified types are valid to be used in FailedMessage
		t.Run("string", func(t *testing.T) {
			strVal := ""
			r := rule.NewSectionExistsRule("type", &strVal, strVal)
			r.Validate()
			assert.NotPanics(t, func() { r.FailedMessage() })
		})

		t.Run("bool", func(t *testing.T) {
			boolVal := false
			r := rule.NewSectionExistsRule("type", &boolVal, boolVal)
			r.Validate()
			assert.NotPanics(t, func() { r.FailedMessage() })
		})

		t.Run("boolString bool", func(t *testing.T) {
			boolStringBoolVal := composerCustomType.BoolString{BoolVal: false, IsBool: true}
			r := rule.NewSectionExistsRule("type", &boolStringBoolVal, boolStringBoolVal)
			r.Validate()
			assert.NotPanics(t, func() { r.FailedMessage() })
		})

		t.Run("boolString string", func(t *testing.T) {
			boolStringStringVal := composerCustomType.BoolString{BoolVal: false, IsBool: true}
			r := rule.NewSectionExistsRule("type", &boolStringStringVal, boolStringStringVal)
			r.Validate()
			assert.NotPanics(t, func() { r.FailedMessage() })
		})
	})
}
