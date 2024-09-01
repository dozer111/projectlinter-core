package config

import "github.com/Masterminds/semver/v3"

type NPMDependency struct {
	name       string
	version    *semver.Version
	versionRaw string
}

func NewNPMDependency(
	name, versionRaw string,
	version *semver.Version,
) *NPMDependency {
	return &NPMDependency{
		name,
		version,
		versionRaw,
	}
}

func (d NPMDependency) Name() string {
	return d.name
}

func (d NPMDependency) Version() *semver.Version {
	return d.version
}

func (d NPMDependency) VersionRaw() string {
	return d.versionRaw
}
