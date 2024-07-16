package utilJSON

import (
	"github.com/nsf/jsondiff"
)

func JSONsAreEqual(js1, js2 []byte) bool {
	opts := jsondiff.DefaultJSONOptions()
	opts.PrintTypes = false
	opts.SkipMatches = false
	opts.Indent = ""

	result, _ := jsondiff.Compare(js1, js2, &opts)
	return result == jsondiff.FullMatch
}
