package config

import (
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_lock"
	"regexp"
	"strconv"

	"github.com/Masterminds/semver/v3"
)

type Config struct {
	Name             string
	PHP              *PHP
	Description      *string
	MinimumStability *string
	PreferStable     *bool
	Type             *string
	Licence          *string
	Dependencies     *composer_json.ComposerDependencies
	Config           *composer_json.RawComposerJsonConfigSection
	Scripts          *composer_json.Scripts
}

type PHP struct {
	Constraint *string
	Platform   *string
}

// in ^8.2
// out 8.2
func (p *PHP) TrimConstraint() string {
	if p.Constraint == nil {
		return ""
	}

	v := *p.Constraint

	re := regexp.MustCompile("[^0-9.]")
	output := re.ReplaceAllString(v, "")

	return output
}

func (p *PHP) PlatformExists() bool {
	return nil != p.Platform
}

func (p *PHP) PlatformToFloat64() float64 {
	platform := *p.Platform
	phpVersion, _ := strconv.ParseFloat(platform, 64)
	return phpVersion
}

func NewComposerConfig(
	composerJson *composer_json.RawComposerJson,
	composerLock *composer_lock.RawComposerLock,
) *Config {
	dependencies := composer_json.NewComposerDependencies()
	devDependencies := composer_json.NewComposerDependencies()

	composerLockPackages := composer_lock.NewComposerLockPackages(composerLock.Packages)
	composerLockPackagesDev := composer_lock.NewComposerLockPackages(composerLock.PackagesDev)
	composerLockAliases := composer_lock.NewComposerLockAliases(composerLock.Aliases)

	for library, constraint := range composerJson.Require {

		var libraryVersion *semver.Version
		if true == composerLockPackages.Has(library) {
			composerLockLibraryVersion, err := semver.NewVersion(composerLockPackages.Get(library).Version)
			if err != nil && composerLockAliases.Has(library) {
				composerLockLibraryVersion, _ = semver.NewVersion(composerLockAliases.Get(library).Alias)
			}

			libraryVersion = composerLockLibraryVersion
		}

		dependencies.Add(composer_json.NewComposerDependency(
			library,
			constraint,
			libraryVersion,
		))
	}

	for library, constraint := range composerJson.RequireDev {

		var libraryVersion *semver.Version
		if true == composerLockPackagesDev.Has(library) {
			composerLockLibraryVersion, err := semver.NewVersion(composerLockPackagesDev.Get(library).Version)
			if err != nil && composerLockAliases.Has(library) {
				composerLockLibraryVersion, _ = semver.NewVersion(composerLockAliases.Get(library).Alias)
			}

			libraryVersion = composerLockLibraryVersion
		}

		devDependencies.Add(composer_json.NewComposerDependency(
			library,
			constraint,
			libraryVersion,
		))
	}

	var composerConstraint *string
	if true == dependencies.Has("php") {
		v := dependencies.Get("php").VersionRaw()
		composerConstraint = &v
	}

	platform, ok := composerLock.PlatformOverrides["php"]
	var composerPlatform *string
	if true == ok {
		composerPlatform = &platform
	}

	var scripts *composer_json.Scripts
	if composerJson.Scripts != nil {
		scripts = composer_json.NewScripts(
			composerJson.Scripts.Arrays,
			composerJson.Scripts.Objects,
			composerJson.Scripts.Strings,
		)
	}

	return &Config{
		composerJson.Name,
		&PHP{
			Constraint: composerConstraint,
			Platform:   composerPlatform,
		},
		composerJson.Description,
		composerJson.MinimumStability,
		composerJson.PreferStable,
		composerJson.Type,
		composerJson.Licence,
		dependencies.Merge(devDependencies),
		composerJson.Config,
		scripts,
	}
}
