package package_json

// RawPackageJSON
// json-schema: https://json.schemastore.org/package.json
type RawPackageJSON struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}
