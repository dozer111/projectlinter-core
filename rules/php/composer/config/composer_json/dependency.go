package composer_json

import (
	"github.com/Masterminds/semver/v3"
)

type ComposerDependency struct {
	name       string
	version    *semver.Version
	constraint string
}

func NewComposerDependency(
	name,
	constraint string,
	version *semver.Version,
) *ComposerDependency {
	return &ComposerDependency{
		name:       name,
		version:    version,
		constraint: constraint,
	}
}

func (d *ComposerDependency) Name() string {
	return d.name
}

func (d *ComposerDependency) Version() *semver.Version {
	return d.version
}

func (d *ComposerDependency) VersionIsCorrect() bool {
	return d.version != nil
}

func (d *ComposerDependency) VersionRaw() string {
	return d.constraint
}

func (d *ComposerDependency) Constraint() string {
	return d.constraint
}
