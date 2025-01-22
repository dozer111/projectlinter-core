package rule

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/dozer111/projectlinter-core/printer"
	"github.com/dozer111/projectlinter-core/rules"

	composerCustomType "github.com/dozer111/projectlinter-core/rules/php/composer/config/composer_json/type"
)

// SectionIsAbsentRule checking that some "simple" block(the value is primitive) is missing
type SectionIsAbsentRule struct {
	val      any
	section  []string
	isPassed bool
}

var _ rules.Rule = (*SectionIsAbsentRule)(nil)

func NewSectionIsAbsentRule[T comparable](section string, val *T) *SectionIsAbsentRule {
	return &SectionIsAbsentRule{
		val:     val,
		section: strings.Split(section, ":"),
	}
}

func (r *SectionIsAbsentRule) ID() string {
	return fmt.Sprintf("composer.%s.absent", strings.Join(r.section, "."))
}

func (r *SectionIsAbsentRule) Title() string {
	return fmt.Sprintf(`"%s" is absent`, strings.Join(r.section, "."))
}

func (r *SectionIsAbsentRule) Validate() {
	// https://codefibershq.com/blog/golang-why-nil-is-not-always-nil
	dynamicValueIsNil := (*[2]uintptr)(unsafe.Pointer(&r.val))[1] == 0
	r.isPassed = dynamicValueIsNil
}

func (r *SectionIsAbsentRule) IsPassed() bool {
	return r.isPassed
}

func (r *SectionIsAbsentRule) FailedMessage() []string {
	if len(r.section) == 1 {
		return printer.NewCodeHintPrinter(
			nil,
			[]string{r.failedMessageNewCode(r.section[0], r.val)},
			nil,
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
			newCode = []string{fmt.Sprintf(`%s%s`, tab, r.failedMessageNewCode(s, r.val))}
		}
	}

	return printer.NewCodeHintPrinter(
		codeBefore,
		newCode,
		nil,
		nil,
	).Print()
}

func (r *SectionIsAbsentRule) failedMessageNewCode(section string, value any) string {
	switch value.(type) {
	case *string:
		return fmt.Sprintf(`"%s": "%s",`, section, *value.(*string))
	case *bool:
		return fmt.Sprintf(`"%s": %t,`, section, *value.(*bool))
	case *composerCustomType.BoolString:
		v := *value.(*composerCustomType.BoolString)
		if v.IsBool {
			return r.failedMessageNewCode(section, &v.BoolVal)
		}

		return r.failedMessageNewCode(section, &v.StrVal)
	default:
		panic(fmt.Sprintf("type %T is not supported", r.val))
	}
}
