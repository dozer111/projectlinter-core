package rule_test

import (
	"github.com/dozer111/projectlinter-core/rules/php/composer/rule"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSectionIsAbsentPassedWhenSectionIsAbsent(t *testing.T) {
	var str *string

	cases := []struct {
		description string
		value       *string
	}{
		{
			"nil",
			nil,
		},
		{
			"var val *string",
			str,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.description, func(t *testing.T) {
			r := rule.NewSectionIsAbsentRule("type", testCase.value)
			r.Validate()

			assert.True(t, r.IsPassed())
		})
	}
}

func TestSectionIsAbsentFailedWhenSectionExists(t *testing.T) {
	stringToPointer := func(val string) *string {
		return &val
	}

	cases := []struct {
		description string
		value       *string
	}{
		{
			"empty string",
			stringToPointer(""),
		},
		{
			"not-empty string",
			stringToPointer("vailkau"),
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.description, func(t *testing.T) {
			r := rule.NewSectionIsAbsentRule("type", testCase.value)
			r.Validate()

			assert.False(t, r.IsPassed())
		})
	}
}
