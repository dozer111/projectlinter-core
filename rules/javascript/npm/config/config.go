package config

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/dozer111/projectlinter-core/rules/javascript/npm/config/package_json"
	"github.com/dozer111/projectlinter-core/rules/javascript/npm/config/package_lock_json"
)

type Config struct {
	Dependencies *NPMDependencies
}

func NewNPMConfig(
	pkg package_json.RawPackageJSON,
	pkgLock package_lock_json.RawPackageLockJSON,
) *Config {
	dependenciesLen := len(pkg.Dependencies) + len(pkg.DevDependencies)
	dependencies := NewNPMDependencies(dependenciesLen)
	devDependencies := NewNPMDependencies(dependenciesLen)

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
