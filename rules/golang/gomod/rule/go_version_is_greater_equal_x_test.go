package rule

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
)

func TestGOVersionIsGreaterEqualRulePass(t *testing.T) {
	cases := []struct {
		description    string
		currentVersion *semver.Version
		minVersion     *semver.Version
	}{
		{
			"currentVersion == expected",
			semver.MustParse("3.15.12"),
			semver.MustParse("3.15.12"),
		},
		{
			"currentVersion > expected",
			semver.MustParse("3.15.12"),
			semver.MustParse("2"),
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.description, func(t *testing.T) {
			r := NewGOVersionIsGreaterEqualRule(*testCase.currentVersion, *testCase.minVersion)
			r.Validate()

			assert.True(t, r.IsPassed())
		})
	}
}

func TestGOVersionIsGreaterEqualRuleFail(t *testing.T) {
	cases := []struct {
		description    string
		currentVersion *semver.Version
		minVersion     *semver.Version
	}{
		{
			"currentVersion < expected",
			semver.MustParse("2"),
			semver.MustParse("4.1"),
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.description, func(t *testing.T) {
			r := NewGOVersionIsGreaterEqualRule(*testCase.currentVersion, *testCase.minVersion)
			r.Validate()

			assert.False(t, r.IsPassed())
		})
	}
}
