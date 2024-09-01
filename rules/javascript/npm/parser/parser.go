package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/1set/gut/yos"

	"github.com/dozer111/projectlinter-core/rules/javascript/npm/config/package_json"
	"github.com/dozer111/projectlinter-core/rules/javascript/npm/config/package_lock_json"
)

type Parser struct {
	dirWithConfigs string
}

var (
	PackageJSONNotFound     = errors.New("package.json not found")
	PackageLockJSONNotFound = errors.New("package-lock.json not found")
)

func NewParser(dirWithConfigs string) *Parser {
	return &Parser{
		dirWithConfigs,
	}
}

func (p *Parser) Parse() (*package_json.RawPackageJSON, *package_lock_json.RawPackageLockJSON, error) {
	packageJSONPath := fmt.Sprintf("%s/package.json", p.dirWithConfigs)
	if false == yos.ExistFile(packageJSONPath) {
		return nil, nil, fmt.Errorf("%w: %s", PackageJSONNotFound, packageJSONPath)
	}

	packageLockJSONPath := fmt.Sprintf("%s/package-lock.json", p.dirWithConfigs)
	if false == yos.ExistFile(packageLockJSONPath) {
		return nil, nil, fmt.Errorf("%w: %s", PackageLockJSONNotFound, packageLockJSONPath)
	}

	packageJson, err := p.parsePackageJSON(packageJSONPath)
	if err != nil {
		return nil, nil, err
	}

	packageLock, err := p.parsePackageLockJSON(packageLockJSONPath)
	if err != nil {
		return nil, nil, err
	}

	return packageJson, packageLock, nil
}

func (p *Parser) parsePackageJSON(packageJSONPath string) (*package_json.RawPackageJSON, error) {
	file, err := os.Open(packageJSONPath)
	if err != nil {
		return nil, fmt.Errorf("cannot open file %s: %w", packageJSONPath, err)
	}
	defer file.Close()

	bytes, _ := io.ReadAll(file)
	var packageJSON package_json.RawPackageJSON
	err = json.Unmarshal(bytes, &packageJSON)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal file to package_json.RawPackageJSON: %w", err)
	}

	return &packageJSON, nil
}

func (p *Parser) parsePackageLockJSON(packageLockPath string) (*package_lock_json.RawPackageLockJSON, error) {
	file, err := os.Open(packageLockPath)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("cannot open file %s: %w", packageLockPath, err)
	}

	bytes, _ := io.ReadAll(file)
	var packageLock package_lock_json.RawPackageLockJSON
	err = json.Unmarshal(bytes, &packageLock)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal file to package_lock_json.RawPackageLockJSON: %w", err)
	}

	return &packageLock, nil
}
