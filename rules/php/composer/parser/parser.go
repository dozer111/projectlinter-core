package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_lock"
	"io"
	"os"

	"github.com/1set/gut/yos"
)

type Parser struct {
	dirWithConfigs string
}

var (
	ComposerJsonNotFound = errors.New("composer.json not found")
	ComposerLockNotFound = errors.New("composer.lock not found")
)

func NewParser(dirWithConfigs string) *Parser {
	return &Parser{
		dirWithConfigs,
	}
}

func (p *Parser) Parse() (*composer_json.RawComposerJson, *composer_lock.RawComposerLock, error) {
	composerJsonPath := fmt.Sprintf("%s/composer.json", p.dirWithConfigs)
	if false == yos.ExistFile(composerJsonPath) {
		return nil, nil, fmt.Errorf("%w: %s", ComposerJsonNotFound, composerJsonPath)
	}

	composerLockPath := fmt.Sprintf("%s/composer.lock", p.dirWithConfigs)
	if false == yos.ExistFile(composerLockPath) {
		return nil, nil, fmt.Errorf("%w: %s", ComposerLockNotFound, composerLockPath)
	}

	composerJson, err := p.parseComposerJson(composerJsonPath)
	if err != nil {
		return nil, nil, err
	}

	composerLock, err := p.parseComposerLock(composerLockPath)
	if err != nil {
		return nil, nil, err
	}

	return composerJson, composerLock, nil
}

func (p *Parser) parseComposerJson(composerJsonPath string) (*composer_json.RawComposerJson, error) {
	file, err := os.Open(composerJsonPath)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("cannot open file %s: %w", composerJsonPath, err)
	}

	bytes, _ := io.ReadAll(file)
	var composerJson composer_json.RawComposerJson
	err = json.Unmarshal(bytes, &composerJson)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal file to composer_json.RawComposerJson: %w", err)
	}

	return &composerJson, nil
}

func (p *Parser) parseComposerLock(composerLockPath string) (*composer_lock.RawComposerLock, error) {
	file, err := os.Open(composerLockPath)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("cannot open file %s: %w", composerLockPath, err)
	}

	bytes, _ := io.ReadAll(file)
	var composerLock composer_lock.RawComposerLock
	err = json.Unmarshal(bytes, &composerLock)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal file to composer_lock.RawComposerLock: %w", err)
	}

	return &composerLock, nil
}
