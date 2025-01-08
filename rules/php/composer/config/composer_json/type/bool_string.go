package composerCustomType

import (
	"encoding/json"
	"errors"
)

// BoolString can hold either a string or a bool.
type BoolString struct {
	IsBool  bool
	BoolVal bool
	StrVal  string
}

var _ json.Unmarshaler = (*BoolString)(nil)

func (sb *BoolString) UnmarshalJSON(data []byte) error {
	var boolVal bool
	if err := json.Unmarshal(data, &boolVal); err == nil {
		sb.IsBool = true
		sb.BoolVal = boolVal
		return nil
	}

	var stringVal string
	if err := json.Unmarshal(data, &stringVal); err == nil {
		sb.IsBool = false
		sb.StrVal = stringVal
		return nil
	}

	return errors.New("value is neither bool nor string")
}
