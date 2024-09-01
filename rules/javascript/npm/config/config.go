package config

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/dozer111/projectlinter-core/rules/javascript/npm/config/package_json"
	"github.com/dozer111/projectlinter-core/rules/javascript/npm/config/package_lock_json"
)

type Config struct {
	Dependencies *package_json.NPMDependencies
}

func NewNPMConfig(
	pkg package_json.RawPackageJSON,
	pkgLock package_lock_json.RawPackageLockJSON,
) *Config {
	dependenciesLen := len(pkg.Dependencies) + len(pkg.DevDependencies)
	dependencies := package_json.NewNPMDependencies(dependenciesLen)
	devDependencies := package_json.NewNPMDependencies(dependenciesLen)

	packageLockPackages := package_lock_json.NewNPMLockPackages(pkgLock.Packages)

	for library, constraint := range pkg.Dependencies {
		var libraryVersion *semver.Version
		libraryPackageName := fmt.Sprintf("node_modules/%s", library)
		if true == packageLockPackages.Has(libraryPackageName) {
			libraryVersion, _ = semver.NewVersion(packageLockPackages.Get(library).Version)
		}

		dependencies.Add(NewNPMDependency(
			library,
			constraint,
			libraryVersion,
		))
	}

	for library, constraint := range pkg.DevDependencies {
		var libraryVersion *semver.Version
		libraryPackageName := fmt.Sprintf("node_modules/%s", library)
		if true == packageLockPackages.Has(libraryPackageName) {
			libraryVersion, _ = semver.NewVersion(packageLockPackages.Get(library).Version)
		}

		devDependencies.Add(NewNPMDependency(
			library,
			constraint,
			libraryVersion,
		))
	}

	return &Config{
		dependencies.Merge(devDependencies),
	}
}

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
