package composer_json

import (
	"encoding/json"
	"fmt"
	"reflect"

	composerCustomType "github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json/type"
)

type RawComposerJson struct {
	Name             string                          `json:"name"`
	Type             *string                         `json:"type,omitempty"`
	Description      *string                         `json:"description,omitempty"`
	Licence          *string                         `json:"license,omitempty"`
	PreferStable     *bool                           `json:"prefer-stable,omitempty"`
	MinimumStability *string                         `json:"minimum-stability,omitempty"`
	Require          map[string]string               `json:"require"`
	RequireDev       map[string]string               `json:"require-dev"`
	Autoload         *RawComposerJsonAutoloadSection `json:"autoload"`
	AutoloadDev      *RawComposerJsonAutoloadSection `json:"autoload-dev"`
	Config           *RawComposerJsonConfigSection   `json:"config"`
	Conflict         map[string]string               `json:"conflict"`
	Scripts          *RawScripts                     `json:"scripts,omitempty"`
}

type RawComposerJsonConfigSection struct {
	SortPackages    *bool                          `json:"sort-packages,omitempty"`
	BumpAfterUpdate *composerCustomType.BoolString `json:"bump-after-update,omitempty"`
	Platform        *map[string]string             `json:"platform,omitempty"`
	AllowPlugins    *map[string]bool               `json:"allow-plugins,omitempty"`
}

type RawComposerJsonAutoloadSection struct {
	Psr4 map[string]string `json:"psr-4"`
}

type RawScripts struct {
	Arrays  map[string][]string
	Objects map[string]map[string]string
	Strings map[string]string
}

func NewRawScriptsFromScripts(scripts *Scripts) *RawScripts {
	return &RawScripts{
		Arrays:  scripts.Arrays,
		Objects: scripts.Objects,
		Strings: scripts.Strings,
	}
}

func (s *RawScripts) UnmarshalJSON(data []byte) error {
	// Temporary map to unmarshal the raw JSON data
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Initialize the maps
	s.Arrays = make(map[string][]string)
	s.Objects = make(map[string]map[string]string)
	s.Strings = make(map[string]string)

	// Iterate over the key-value pairs in the raw map
	for key, value := range raw {
		vType := reflect.TypeOf(value).Kind()
		switch vType {
		case reflect.Slice:
			// Handle slice (array) type
			arr := value.([]interface{})
			strArr := make([]string, len(arr))
			for i, v := range arr {
				strArr[i] = v.(string)
			}
			s.Arrays[key] = strArr
		case reflect.Map:
			// Handle map (object) type
			obj := value.(map[string]interface{})
			strObj := make(map[string]string)
			for k, v := range obj {
				strObj[k] = v.(string)
			}
			s.Objects[key] = strObj
		case reflect.String:
			// Handle string type
			s.Strings[key] = value.(string)
		default:
			fmt.Printf("Key: %s has an unhandled type\n", key)
		}
	}

	return nil
}

func (s *RawScripts) MarshalJSON() ([]byte, error) {
	// Temporary map to hold the serialized data
	raw := make(map[string]interface{})

	// Populate the map with the contents of Arrays, Objects, and Strings
	for key, value := range s.Arrays {
		raw[key] = value
	}
	for key, value := range s.Objects {
		raw[key] = value
	}
	for key, value := range s.Strings {
		raw[key] = value
	}

	// Marshal the map to JSON
	return json.Marshal(raw)
}
