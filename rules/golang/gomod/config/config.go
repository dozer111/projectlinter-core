package config

import (
	"github.com/Masterminds/semver/v3"
	"github.com/rogpeppe/go-internal/modfile"
)

type Config struct {
	Modfile      *modfile.File
	Dependencies *GomodDependencies
}

type GomodDependency struct {
	name       string
	version    *semver.Version
	versionRaw string
}

func NewGomodDependency(
	name, versionRaw string,
	version *semver.Version,
) *GomodDependency {
	return &GomodDependency{
		name,
		version,
		versionRaw,
	}
}

func (d GomodDependency) Name() string {
	return d.name
}

func (d GomodDependency) Version() *semver.Version {
	return d.version
}

func (d GomodDependency) VersionRaw() string {
	return d.versionRaw
}
