package rule

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dozer111/projectlinter-core/printer"
	"github.com/dozer111/projectlinter-core/rules"
)

// SectionHasCorrectValueRule check the simple section(the value is primitive) value is correct
type SectionHasCorrectValueRule[T comparable] struct {
	value    sectionValue[T]
	section  []string
	isPassed bool
}

var _ rules.Rule = (*SectionHasCorrectValueRule[string])(nil)

type sectionValue[T comparable] struct {
	expected []T
	actual   T
}

func NewSectionHasCorrectValueRule[T comparable](
	section string,
	actualValue T,
	expectedValue ...T,
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
	if len(r.value.expected) == 0 {
		r.isPassed = false
		return
	}

	for _, v := range r.value.expected {
		if r.value.actual == v {
			r.isPassed = true
			break
		}
	}
}

func (r *SectionHasCorrectValueRule[T]) IsPassed() bool {
	return r.isPassed
}

func (r *SectionHasCorrectValueRule[T]) FailedMessage() []string {
	if len(r.value.expected) == 0 {
		return []string{
			"rule must contain at least one expected value",
		}
	}

	resultMessage := make([]string, 0, 4)
	if len(r.value.expected) > 1 {
		resultMessage = append(
			resultMessage,
			"There is a number of correct values for this section.",
			"Choose the most correct one",
		)
	}

	actual := r.value.actual

	if len(r.section) == 1 {
		newCode := make([]string, 0, len(r.value.expected))
		for _, expected := range r.value.expected {
			newCode = append(newCode, r.failedMessageNewCode(r.section[0], expected))
		}

		message := printer.NewCodeHintPrinter(
			nil,
			[]string{r.failedMessageNewCode(r.section[0], actual)},
			newCode,
			nil,
		).Print()

		return slices.Concat(resultMessage, message)
	}

	var codeBefore []string
	var newCode []string
	var wrongCode string
	for i, s := range r.section {
		tab := strings.Repeat("\t", i)
		if i != len(r.section)-1 {
			codeBefore = append(codeBefore, fmt.Sprintf(`%s"%s": {`, tab, s))
		} else {
			codeBefore = append(codeBefore, fmt.Sprintf(`%s"..": ...,`, tab))
			wrongCode = fmt.Sprintf(`%s%s`, tab, r.failedMessageNewCode(s, actual))
			for _, expected := range r.value.expected {
				newCode = append(newCode, r.failedMessageNewCode(s, expected))
			}
		}
	}

	return slices.Concat(
		resultMessage,
		printer.NewCodeHintPrinter(
			codeBefore,
			[]string{wrongCode},
			newCode,
			nil,
		).Print(),
	)
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
