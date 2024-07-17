package utilTest

import "github.com/dozer111/projectlinter-core/rules"

func AllSetRulesArePassed(rules []rules.Rule) []string {
	failedRules := make([]string, 0, len(rules))

	for _, r := range rules {
		if !r.IsPassed() {
			failedRules = append(failedRules, r.ID())
		}
	}

	return failedRules
}
