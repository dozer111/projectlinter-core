package rule

import (
	"fmt"
	"github.com/dozer111/projectlinter-core/printer"
	"github.com/dozer111/projectlinter-core/rules"
	"strings"
)

// перевірка що значення "простого"(значення якого - примітивний тип даних) блоку вірне
type SectionHasCorrectValueRule[T comparable] struct {
	value    sectionValue[T]
	section  []string
	isPassed bool
}

var _ rules.Rule = (*SectionHasCorrectValueRule[string])(nil)

type sectionValue[T comparable] struct {
	expected T
	actual   T
}

func (v sectionValue[T]) defaultActual() T {
	var defaultValue T
	return defaultValue
}

func NewSectionHasCorrectValueRule[T comparable](
	section string,
	expectedValue T,
	actualValue T,
) *SectionHasCorrectValueRule[T] {
	return &SectionHasCorrectValueRule[T]{
		value: sectionValue[T]{
			expectedValue,
			actualValue,
		},
		section: strings.Split(section, ":"),
	}
}

func (r *SectionHasCorrectValueRule[T]) ID() string {
	return fmt.Sprintf("composer.%s.valid", strings.Join(r.section, "."))
}

func (r *SectionHasCorrectValueRule[T]) Title() string {
	return fmt.Sprintf(`"%s" has correct value`, strings.Join(r.section, "."))
}

func (r *SectionHasCorrectValueRule[T]) Validate() {
	r.isPassed = r.value.actual == r.value.expected
}

func (r *SectionHasCorrectValueRule[T]) IsPassed() bool {
	return r.isPassed
}

func (r *SectionHasCorrectValueRule[T]) FailedMessage() []string {
	actual := r.value.actual

	if len(r.section) == 1 {
		return printer.NewCodeHintPrinter(
			nil,
			[]string{r.failedMessageNewCode(r.section[0], actual)},
			[]string{r.failedMessageNewCode(r.section[0], r.value.expected)},
			nil,
		).Print()
	}

	var codeBefore []string
	var wrongCode, newCode string
	for i, s := range r.section {
		tab := strings.Repeat("\t", i)
		if i != len(r.section)-1 {
			codeBefore = append(codeBefore, fmt.Sprintf(`%s"%s": {`, tab, s))
		} else {
			codeBefore = append(codeBefore, fmt.Sprintf(`%s"..": ...,`, tab))
			wrongCode = fmt.Sprintf(`%s%s`, tab, r.failedMessageNewCode(s, actual))
			newCode = fmt.Sprintf(`%s%s`, tab, r.failedMessageNewCode(s, r.value.expected))
		}
	}

	return printer.NewCodeHintPrinter(
		codeBefore,
		[]string{wrongCode},
		[]string{newCode},
		nil,
	).Print()
}

func (r *SectionHasCorrectValueRule[T]) failedMessageNewCode(section string, value any) string {
	switch value.(type) {
	case string:
		return fmt.Sprintf(`"%s": "%s",`, section, value)
	case bool:
		return fmt.Sprintf(`"%s": %t,`, section, value)
	default:
		panic(fmt.Sprintf("type %T is not supported", value))
	}
}
