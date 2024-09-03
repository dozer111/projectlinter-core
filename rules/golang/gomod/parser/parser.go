package parser

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/dozer111/projectlinter-core/rules/golang/gomod/config"

	"github.com/1set/gut/yos"
	"github.com/Masterminds/semver/v3"
	"github.com/rogpeppe/go-internal/modfile"
)

type Parser struct {
	pathToDirWithGoMod string
}

func NewParser(pathToDirWithGoMod string) *Parser {
	return &Parser{pathToDirWithGoMod}
}

var (
	GoModIsAbsent = errors.New("go.mod is absent")
)

func (p *Parser) Parse() (*config.Config, error) {
	pathToGoMod := fmt.Sprintf("%s/go.mod", p.pathToDirWithGoMod)
	if !yos.ExistFile(pathToGoMod) {
		return nil, fmt.Errorf("%w: %s", GoModIsAbsent, p.pathToDirWithGoMod)
	}

	file, err := os.Open(pathToGoMod)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("cannot open file %s: %w", pathToGoMod, err)
	}

	gomod, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %s", pathToGoMod)
	}

	var mf *modfile.File
	mf, err = modfile.Parse("", gomod, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot parse modfile: %w", err)
	}

	result := &config.Config{
		Modfile:      mf,
		Dependencies: config.NewGODependencies(len(mf.Require)),
	}

	for _, d := range mf.Require {
		v, _ := semver.NewVersion(d.Mod.Version)
		result.Dependencies.Add(config.NewGomodDependency(d.Mod.Path, d.Mod.Version, v, d.Indirect))
	}

	return result, nil
}
