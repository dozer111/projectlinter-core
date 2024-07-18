package substitute

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
	PathToBumpConfigsIsInvalid = errors.New("path to bump configs is invalid")
	BumpConfigDoesNotFitSchema = errors.New("bump config does not fit schema")
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
		return nil, fmt.Errorf("%w: %s", PathToBumpConfigsIsInvalid, p.pathToConfigs)
	}

	librariesToChange := make([]Library, 0, 20)

	configs, _ := yos.ListFile(p.pathToConfigs)
	for _, configFile := range configs {
		if configFile.Info.IsDir() ||
			!strings.HasSuffix(configFile.Info.Name(), ".yaml") {
			continue
		}

		data, err := os.ReadFile(configFile.Path)

		if err != nil {
			return nil, fmt.Errorf("cannot open file %s: %w", configFile.Path, err)
		}

		if err := p.assertConfigSatisfiesToJsonSchema(data, configFile.Path); err != nil {
			return nil, err
		}

		var parsedConfig Library
		if err = yaml.Unmarshal(data, &parsedConfig); err != nil {
			return nil, fmt.Errorf("cannot unmarshal yaml file %s: %w", configFile.Path, err)
		}

		librariesToChange = append(librariesToChange, parsedConfig)
	}

	return librariesToChange, nil
}

func (p *Parser) assertConfigSatisfiesToJsonSchema(data []byte, filePath string) error {
	schemaLoader := gojsonschema.NewStringLoader(jsonSchema)

	var rawConfig interface{}
	err := yaml.Unmarshal(data, &rawConfig)

	if err != nil {
		return fmt.Errorf("cannot unmarshal data from file %s to interface{}: %w", filePath, err)
	}
	documentLoader := gojsonschema.NewRawLoader(rawConfig)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return fmt.Errorf("cannot ensure that file %s fits substitute json-schema: %w", filePath, err)
	}
	if result.Valid() == false {
		return fmt.Errorf("%w: %s: %v", BumpConfigDoesNotFitSchema, filePath, result.Errors())
	}

	return nil
}
