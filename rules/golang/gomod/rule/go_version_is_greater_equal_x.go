package rule

import (
	"fmt"
	"github.com/dozer111/projectlinter-core/rules"

	"github.com/Masterminds/semver/v3"
)

type GOVersionIsGreaterEqual struct {
	currentVersion  semver.Version
	expectedVersion semver.Version

	isPassed bool
}

var _ rules.Rule = (*GOVersionIsGreaterEqual)(nil)

func NewGOVersionIsGreaterEqualRule(
	currentVersion,
	expectedVersion semver.Version,
) *GOVersionIsGreaterEqual {
	return &GOVersionIsGreaterEqual{
		currentVersion:  currentVersion,
		expectedVersion: expectedVersion,
	}
}

func (r *GOVersionIsGreaterEqual) ID() string {
	return "go_version_is_correct"
}

func (r *GOVersionIsGreaterEqual) Title() string {
	return fmt.Sprintf("go version >= %s", r.expectedVersion.String())
}

func (r *GOVersionIsGreaterEqual) Validate() {
	r.isPassed = !r.currentVersion.LessThan(&r.expectedVersion)
}

func (r *GOVersionIsGreaterEqual) IsPassed() bool {
	return r.isPassed
}

func (r *GOVersionIsGreaterEqual) FailedMessage() []string {
	return []string{
		fmt.Sprintf("go must be at least %s", r.expectedVersion.String()),
	}
}
