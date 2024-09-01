package package_lock_json

type RawPackageLockJSON struct {
	Packages map[string]*RawPackageLockPackage `json:"packages"`
}

type RawPackageLockPackage struct {
	Version string `json:"version"`
}
