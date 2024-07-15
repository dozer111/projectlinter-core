package painter_test

import (
	"github.com/dozer111/projectlinter/util/painter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBashPainter(t *testing.T) {
	p := painter.NewPainter()

	var testCases = []struct {
		description    string
		expectedOutput string
		actualOutput   string
	}{
		{
			"Red",
			"\x1b[31msome text\x1b[0m",
			p.Red("some text"),
		},
		{
			"Green",
			"\x1b[32msome text\x1b[0m",
			p.Green("some text"),
		},
		{
			"Blue",
			"\x1b[34msome text\x1b[0m",
			p.Blue("some text"),
		},
		{
			"White",
			"\x1b[37msome text\x1b[0m",
			p.White("some text"),
		},
		{
			"Yellow",
			"\x1b[33msome text\x1b[0m",
			p.Yellow("some text"),
		},
		{
			"Warning",
			"\x1b[33m⚠ some text\x1b[0m",
			p.Warning("some text"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.expectedOutput, tc.actualOutput)
		})
	}
}
