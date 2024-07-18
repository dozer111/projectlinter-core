package bump

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/1set/gut/yos"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

var (
	PathToConfigsIsInvalid = errors.New("path to substitute configs is invalid")
	ConfigDoesNotFitSchema = errors.New("substitute config does not fit schema")
)

//go:embed json-schema.json
var jsonSchema string

type Parser struct {
	pathToConfigs string
}

func NewParser(pathToConfigs string) *Parser {
	return &Parser{pathToConfigs}
}

func (p *Parser) Parse() ([]Library, error) {
	if !yos.Exist(p.pathToConfigs) {
		return nil, fmt.Errorf("%w: %s", PathToConfigsIsInvalid, p.pathToConfigs)
	}

	var bumpConfigs []Library

	configs, _ := yos.ListFile(p.pathToConfigs)
	for _, realConfigFile := range configs {
		if !strings.HasSuffix(realConfigFile.Info.Name(), ".yaml") {
			continue
		}

		data, err := os.ReadFile(realConfigFile.Path)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("cannot open file %s: %v", realConfigFile.Path, err))
		}

		if err := p.assertConfigSatisfiesToJsonSchema(data, realConfigFile.Path); err != nil {
			return nil, err
		}

		var parsedConfig Library
		if err = yaml.Unmarshal(data, &parsedConfig); err != nil {
			return nil, errors.New(fmt.Sprintf("cannot unmarshal yaml %s", realConfigFile.Path))
		}

		bumpConfigs = append(bumpConfigs, parsedConfig)
	}

	return bumpConfigs, nil
}

func (p *Parser) assertConfigSatisfiesToJsonSchema(data []byte, filePath string) error {
	schemaLoader := gojsonschema.NewStringLoader(jsonSchema)

	var rawConfig interface{}
	err := yaml.Unmarshal(data, &rawConfig)
	if err != nil {
		return errors.New(fmt.Sprintf("cannot parse data from %s to interface{}: %v", filePath, err))
	}
	documentLoader := gojsonschema.NewRawLoader(rawConfig)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return errors.New(fmt.Sprintf("cannot check does the file %s fits json-schema: %v", filePath, err))
	}
	if result.Valid() == false {
		return fmt.Errorf("%w: %s: %v", ConfigDoesNotFitSchema, filePath, result.Errors())
	}

	return nil
}
