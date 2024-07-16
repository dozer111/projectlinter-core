package composer_lock

type RawComposerLock struct {
	Platform          map[string]string         `json:"platform"`
	PlatformOverrides map[string]string         `json:"platform-overrides"`
	Packages          []*RawComposerLockPackage `json:"packages"`
	Aliases           []*RawComposerLockAlias   `json:"aliases"`
	PackagesDev       []*RawComposerLockPackage `json:"packages-dev"`
}

type RawComposerLockAlias struct {
	Package string `json:"package"`
	Alias   string `json:"alias"`
}

type RawComposerLockPackage struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
