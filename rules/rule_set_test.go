package rules_test

import (
	"github.com/dozer111/projectlinter-core/rules"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleTreeResolve(t *testing.T) {
	testCases := []struct {
		description string
		tree        *rules.RuleTree
		ruleAmount  int
	}{
		{
			"only root success leafs",
			rules.NewRuleTree(
				leaf(alwaysSuccessRule{}), // 1
				leaf(alwaysSuccessRule{}), // 2
				leaf(alwaysSuccessRule{}), // 3
				leaf(alwaysSuccessRule{}), // 4
				leaf(alwaysSuccessRule{}), // 5
			),
			5,
		},
		{
			"only root leafs",
			rules.NewRuleTree(
				leaf(alwaysSuccessRule{}), // 1
				leaf(alwaysSuccessRule{}), // 2
				leaf(alwaysSuccessRule{}), // 3
				leaf(alwaysFailedRule{}),  // 4
				leaf(alwaysFailedRule{}),  // 5
			),
			5,
		},
		{
			"only success leafs with children",
			rules.NewRuleTree(
				leaf( // 1
					alwaysSuccessRule{},
					leaf(
						alwaysSuccessRule{}, // 2
						leaf(
							alwaysSuccessRule{},       // 3
							leaf(alwaysSuccessRule{}), // 4
						),
					),
				),
				leaf(alwaysSuccessRule{}), // 5
			),
			5,
		},
		{
			"leafs with children",
			rules.NewRuleTree(
				leaf( // 1
					alwaysSuccessRule{},
					leaf(
						alwaysSuccessRule{}, // 2
						leaf(
							alwaysSuccessRule{},      // 3
							leaf(alwaysFailedRule{}), // 4
						),
					),
				),
				leaf(alwaysFailedRule{}), // 5
			),
			5,
		},
		{
			"leafs with children(failed rule has children)",
			rules.NewRuleTree(
				leaf( // 1
					alwaysSuccessRule{},
					leaf(
						alwaysSuccessRule{}, // 2
						leaf(
							alwaysSuccessRule{}, // 3
							leaf(
								alwaysFailedRule{}, // 4
								// parent leaf is failed, so this rules will not be resolved
								leaf(alwaysSuccessRule{}),
								leaf(alwaysSuccessRule{}),
								leaf(alwaysSuccessRule{}),
							),
						),
					),
				),
				leaf(alwaysFailedRule{}), // 5
			),
			5,
		},
		{
			"optional leaf",
			rules.NewRuleTree(
				leaf(alwaysSuccessRule{}),         // 1
				leaf(alwaysSuccessRule{}),         // 2
				optionalLeaf(alwaysSuccessRule{}), // 3
				optionalLeaf(alwaysSuccessRule{}), // 4
				leaf(alwaysSuccessRule{}),         // 5
			),
			5,
		},
		{
			"optional leaf with children",
			rules.NewRuleTree(
				leaf( // 1
					alwaysSuccessRule{},
					optionalLeaf(
						alwaysSuccessRule{}, // 2
						leaf(
							alwaysSuccessRule{},       // 3
							leaf(alwaysSuccessRule{}), // 4
						),
					),
				),
				leaf(alwaysSuccessRule{}), // 5
			),
			5,
		},
		{
			"optional leaf failed",
			rules.NewRuleTree(
				leaf(alwaysSuccessRule{}),        // 1
				leaf(alwaysSuccessRule{}),        // 2
				optionalLeaf(alwaysFailedRule{}), // failed optional rule does not include to result in resolve
				optionalLeaf(alwaysFailedRule{}), // failed optional rule does not include to result in resolve
				leaf(alwaysSuccessRule{}),        // 5
			),
			3,
		},
		{
			"optional leaf with children failed",
			rules.NewRuleTree(
				leaf( // 1
					alwaysSuccessRule{},
					optionalLeaf(
						alwaysFailedRule{}, // 2
						leaf(
							alwaysSuccessRule{},       // 3
							leaf(alwaysSuccessRule{}), // 4
						),
					),
				),
				leaf(alwaysSuccessRule{}), // 5
			),
			2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			tree := testCase.tree
			resolvedRules := tree.Resolve([]string{})

			assert.Equal(t, testCase.ruleAmount, len(resolvedRules))
		})
	}
}

type alwaysSuccessRule struct {
}

func (a alwaysSuccessRule) ID() string {
	return ""
}

func (a alwaysSuccessRule) Title() string {
	return ""
}

func (a alwaysSuccessRule) Validate() {}

func (a alwaysSuccessRule) IsPassed() bool {
	return true
}

func (a alwaysSuccessRule) FailedMessage() []string {
	return nil
}

type alwaysFailedRule struct {
	alwaysSuccessRule
}

func (a alwaysFailedRule) IsPassed() bool {
	return false
}

func leaf(r rules.Rule, children ...rules.RuleTreeLeaf) rules.RuleTreeLeaf {
	return rules.NewLeaf(r, children...)
}

func optionalLeaf(r rules.Rule, children ...rules.RuleTreeLeaf) rules.RuleTreeLeaf {
	return rules.NewOptionalLeaf(r, children...)
}
