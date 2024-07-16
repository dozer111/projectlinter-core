package composer_json

// Scripts structure is made just like that because
// 1. this block can contain data of various types
// 2. I didn't want to be tied to script names
//
// This is done so that the structure(concrete script) must be abstract
// So that anyone can come and write their rule
type Scripts struct {
	Arrays      map[string][]string
	Objects     map[string]map[string]string
	Strings     map[string]string
	internalMap map[string]ScriptsType
}

type ScriptsType int

const (
	TypeArray  ScriptsType = 1
	TypeObject ScriptsType = 2
	TypeString ScriptsType = 3
)

func NewScripts(
	Arrays map[string][]string,
	Objects map[string]map[string]string,
	Strings map[string]string,
) *Scripts {
	scripts := &Scripts{
		Arrays:  Arrays,
		Objects: Objects,
		Strings: Strings,
	}

	scripts.generateInternalMap()
	return scripts
}

func (s *Scripts) Has(key string) (bool, ScriptsType) {
	scriptType, ok := s.internalMap[key]
	return ok, scriptType
}

func (s *Scripts) Len() int {
	return len(s.internalMap)
}

func (s *Scripts) Merge(s2 *Scripts) *Scripts {
	if s2 == nil {
		return NewScripts(s.Arrays, s.Objects, s.Strings)
	}

	mergedArrays := make(map[string][]string)
	mergedObjects := make(map[string]map[string]string)
	mergedStrings := make(map[string]string)

	for k, v := range s.Arrays {
		mergedArrays[k] = v
	}

	for k, v := range s2.Arrays {
		mergedArrays[k] = v
	}

	for k, v := range s.Objects {
		mergedObjects[k] = v
	}

	for k, v := range s2.Objects {
		mergedObjects[k] = v
	}

	for k, v := range s.Strings {
		mergedStrings[k] = v
	}

	for k, v := range s2.Strings {
		mergedStrings[k] = v
	}

	return NewScripts(mergedArrays, mergedObjects, mergedStrings)
}

func (s *Scripts) generateInternalMap() {
	s.internalMap = make(map[string]ScriptsType, len(s.Arrays)+len(s.Objects)+len(s.Strings))

	for k := range s.Arrays {
		s.internalMap[k] = TypeArray
	}

	for k := range s.Objects {
		s.internalMap[k] = TypeObject
	}

	for k := range s.Strings {
		s.internalMap[k] = TypeString
	}
}
