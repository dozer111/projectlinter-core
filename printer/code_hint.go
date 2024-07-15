package printer

import (
	"github.com/dozer111/projectlinter-core/util/painter"
)

// CodeHintPrinter - Hint for the code.
// Use it when
// 1. you want to show that something was missing, and we want to insert it into a specific place.
// 2. OR, something is wrong in specified place and we want to change it
type CodeHintPrinter struct {
	before    []string
	wrongCode []string
	newCode   []string
	after     []string
}

func NewCodeHintPrinter(
	before,
	wrongCode,
	newCode,
	after []string,
) *CodeHintPrinter {
	return &CodeHintPrinter{
		before,
		wrongCode,
		newCode,
		after,
	}
}

func (p *CodeHintPrinter) Print() []string {
	paint := painter.NewPainter()

	var result []string
	if len(p.before) > 0 {
		for _, beforeLine := range p.before {
			result = append(result, paint.Yellow(beforeLine))
		}
	}

	for _, wrongCodeLine := range p.wrongCode {
		result = append(result, paint.Red(wrongCodeLine))
	}

	for _, newCodeLine := range p.newCode {
		result = append(result, paint.Green(newCodeLine))
	}

	if len(p.after) > 0 {
		for _, afterLine := range p.after {
			result = append(result, paint.Yellow(afterLine))
		}
	}

	return result
}
