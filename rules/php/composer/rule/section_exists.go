package rule

import (
	"encoding/json"
	"fmt"
	"github.com/dozer111/projectlinter-core/printer"
	"github.com/dozer111/projectlinter-core/rules"
	"github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json"
	"strings"
	"unsafe"
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
		newCode := r.failedMessageNewCode(r.section[0])

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
			newCode = []string{fmt.Sprintf(`%s%s`, tab, r.failedMessageNewCode(s))}
		}
	}

	return printer.NewCodeHintPrinter(
		codeBefore,
		nil,
		newCode,
		nil,
	).Print()
}

func (r *SectionExistsRule) failedMessageNewCode(section string) string {
	switch r.proposedValue.(type) {
	case string:
		return fmt.Sprintf(`"%s": "%s",`, section, r.proposedValue)
	case bool:
		return fmt.Sprintf(`"%s": %t,`, section, r.proposedValue)
	case composer_json.RawComposerJsonConfigSection:
		js, _ := json.MarshalIndent(r.proposedValue, "", "	")
		v := fmt.Sprintf(`"%s": %s,`, section, js)
		return v
	}

	panic(fmt.Sprintf("type %T is not supported", r.proposedValue))
}
