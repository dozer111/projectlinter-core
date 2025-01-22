package rule

import (
	"encoding/json"
	"fmt"
	"strings"
	"unsafe"

	"github.com/dozer111/projectlinter-core/printer"
	"github.com/dozer111/projectlinter-core/rules"

	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	composerCustomType "github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json/type"
)

// SectionExistsRule checking whether there is any "simple" block(the value is primitive)
type SectionExistsRule struct {
	val           any
	proposedValue any
	section       []string
	isPassed      bool
}

var _ rules.Rule = (*SectionExistsRule)(nil)

func NewSectionExistsRule[T comparable](section string, val *T, proposedValue T) *SectionExistsRule {
	return &SectionExistsRule{
		val:           val,
		proposedValue: proposedValue,
		section:       strings.Split(section, ":"),
	}
}

func (r *SectionExistsRule) ID() string {
	return fmt.Sprintf("composer.%s.exists", strings.Join(r.section, "."))
}

func (r *SectionExistsRule) Title() string {
	return fmt.Sprintf(`"%s" exists`, strings.Join(r.section, "."))
}

func (r *SectionExistsRule) Validate() {
	// https://codefibershq.com/blog/golang-why-nil-is-not-always-nil
	dynamicValueIsNil := (*[2]uintptr)(unsafe.Pointer(&r.val))[1] == 0
	r.isPassed = !dynamicValueIsNil
}

func (r *SectionExistsRule) IsPassed() bool {
	return r.isPassed
}

func (r *SectionExistsRule) FailedMessage() []string {
	if len(r.section) == 1 {
		newCode := r.failedMessageNewCode(r.section[0], r.proposedValue)

		return printer.NewCodeHintPrinter(
			nil,
			nil,
			strings.Split(newCode, "\n"),
			nil,
		).Print()
	}

	var codeBefore []string
	var newCode []string
	for i, s := range r.section {
		tab := strings.Repeat("\t", i)
		if i != len(r.section)-1 {
			codeBefore = append(codeBefore, fmt.Sprintf(`%s"%s": {`, tab, s))
		} else {
			newCode = []string{fmt.Sprintf(`%s%s`, tab, r.failedMessageNewCode(s, r.proposedValue))}
		}
	}

	return printer.NewCodeHintPrinter(
		codeBefore,
		nil,
		newCode,
		nil,
	).Print()
}

func (r *SectionExistsRule) failedMessageNewCode(section string, value any) string {
	switch value.(type) {
	case string:
		return fmt.Sprintf(`"%s": "%s",`, section, value)
	case bool:
		return fmt.Sprintf(`"%s": %t,`, section, value)
	case composer_json.RawComposerJsonConfigSection:
		js, _ := json.MarshalIndent(value, "", "	")
		v := fmt.Sprintf(`"%s": %s,`, section, js)
		return v
	case composerCustomType.BoolString:
		v := value.(composerCustomType.BoolString)
		if v.IsBool {
			return r.failedMessageNewCode(section, v.BoolVal)
		}

		return r.failedMessageNewCode(section, v.StrVal)
	}

	panic(fmt.Sprintf("type %T is not supported", value))
}
