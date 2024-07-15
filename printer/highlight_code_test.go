package printer_test

import (
	"github.com/dozer111/projectlinter-core/printer"
	"github.com/dozer111/projectlinter-core/util/painter"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestHighlightCodePrinter(t *testing.T) {
	painter.CurrentPainterState = painter.Fake

	t.Run("1. add new item to yaml list", func(t *testing.T) {
		input := []string{
			"composer.json",
			"composer.lock",
			"phpunit.xml.dist",
			"rector.php",
		}

		newElements := []string{"rector.php"}
		oldElements := []string{}

		p := printer.NewHighlightCodePrinter(
			input,
			newElements,
			oldElements,
		)

		expectedOutput := `composer.json
composer.lock
phpunit.xml.dist
(green)rector.php`
		assert.Equal(t, expectedOutput, strings.Join(p.Print(), "\n"))
	})

	t.Run("2. change the yaml list", func(t *testing.T) {
		input := []string{
			"composer.json",
			"composer.lock",
			"phpunit.xml.dist",
			"rector.php",
		}

		newElements := []string{"rector.php"}
		oldElements := []string{"composer.json"}

		p := printer.NewHighlightCodePrinter(
			input,
			newElements,
			oldElements,
		)

		expectedOutput := `(red)composer.json
composer.lock
phpunit.xml.dist
(green)rector.php`
		assert.Equal(t, expectedOutput, strings.Join(p.Print(), "\n"))
	})
}
