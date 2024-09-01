package package_lock_json

type RawPackageLockJSON struct {
	Packages map[string]*RawNPMLockPackage `json:"packages"`
}

type RawNPMLockPackage struct {
	Version string `json:"version"`
}
