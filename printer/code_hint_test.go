package printer_test

import (
	"github.com/dozer111/projectlinter/printer"
	"github.com/dozer111/projectlinter/util/painter"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCodeHintPrinter(t *testing.T) {
	painter.CurrentPainterState = painter.Fake

	t.Run("1. add new code", func(t *testing.T) {
		before := []string{
			`"require": {`,
			"\t...",
		}
		after := []string{"\t...", "}"}
		newCode := []string{"\t\"symfony/validator\": \"6\","}

		p := printer.NewCodeHintPrinter(
			before,
			[]string{},
			newCode,
			after,
		)

		expectedOutput := `(yellow)"require": {
(yellow)	...
(green)	"symfony/validator": "6",
(yellow)	...
(yellow)}`
		assert.Equal(t, expectedOutput, strings.Join(p.Print(), "\n"))
	})

	t.Run("2. change current code to new", func(t *testing.T) {
		before := []string{
			`"require": {`,
			"\t...",
		}
		after := []string{"\t...", "}"}
		oldCode := []string{"\t\"symfony/validator\": \"6\","}
		newCode := []string{"\t\"symfony/validator\": \"^6\","}

		p := printer.NewCodeHintPrinter(
			before,
			oldCode,
			newCode,
			after,
		)

		expectedOutput := `(yellow)"require": {
(yellow)	...
(red)	"symfony/validator": "6",
(green)	"symfony/validator": "^6",
(yellow)	...
(yellow)}`
		assert.Equal(t, expectedOutput, strings.Join(p.Print(), "\n"))
	})

}
