package printer

import (
	"github.com/dozer111/projectlinter/util/painter"
)

// HighlightCodePrinter - The main idea is that we have an output ([]string), for example - a YAML list.
// We want to keep the elements of the output as they are and additionally highlight everything in green/red.
type HighlightCodePrinter struct {
	input    []string
	newItems []string
	oldItems []string
}

func NewHighlightCodePrinter(input, new, old []string) *HighlightCodePrinter {
	return &HighlightCodePrinter{
		input,
		new,
		old,
	}
}

func (p *HighlightCodePrinter) Print() []string {
	paint := painter.NewPainter()

	result := make([]string, 0, len(p.input))
	for _, item := range p.input {
		result = append(result, p.print(paint, item))
	}

	return result
}

func (p *HighlightCodePrinter) print(paint painter.Painter, item string) string {
	for _, oldItem := range p.oldItems {
		if item == oldItem {
			return paint.Red(item)
		}
	}

	for _, newItem := range p.newItems {
		if item == newItem {
			return paint.Green(item)
		}
	}

	return item
}
