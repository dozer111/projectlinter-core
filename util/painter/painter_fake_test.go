package painter_test

import (
	"github.com/dozer111/projectlinter/util/painter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFakePainter(t *testing.T) {
	painter.CurrentPainterState = painter.Fake
	p := painter.NewPainter()

	var testCases = []struct {
		description    string
		expectedOutput string
		actualOutput   string
	}{
		{
			"Red",
			"(red)some text",
			p.Red("some text"),
		},
		{
			"Green",
			"(green)some text",
			p.Green("some text"),
		},
		{
			"Blue",
			"(blue)some text",
			p.Blue("some text"),
		},
		{
			"White",
			"(white)some text",
			p.White("some text"),
		},
		{
			"Yellow",
			"(yellow)some text",
			p.Yellow("some text"),
		},
		{
			"Warning",
			"(warn)some text",
			p.Warning("some text"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedOutput, tc.actualOutput)
		})
	}
}
