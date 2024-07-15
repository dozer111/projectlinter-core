package painter

import "fmt"

// bashPainter - colors in bash CLI
type bashPainter struct {
}

var _ Painter = (*bashPainter)(nil)

func (p bashPainter) Red(text string) string {
	return p.paint("\033[31m", text)
}

func (p bashPainter) Green(text string) string {
	return p.paint("\033[32m", text)
}

func (p bashPainter) Blue(text string) string {
	return p.paint("\033[34m", text)
}

func (p bashPainter) White(text string) string {
	return p.paint("\033[37m", text)
}

func (p bashPainter) Yellow(text string) string {
	return p.paint("\033[33m", text)
}

func (p bashPainter) Warning(text string) string {
	return p.Yellow(fmt.Sprintf("⚠ %s", text))
}

func (p bashPainter) paint(color, text string) string {
	return fmt.Sprintf("%s%s%s", color, text, "\033[0m")
}
