package bump

import (
	"github.com/dozer111/projectlinter-core/util/painter"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrinterOnFullCorrectExample(t *testing.T) {
	painter.CurrentPainterState = painter.Fake

	// ARRANGE: some dependency need to be upgraded
	printer := bumpDependencyPrinter{}
	bumpConfig := &bumpDependencyConfig{
		Library: Library{
			"some-private-git/limit-lb",
			"5.0",
			[]string{
				"The internal code was splitted to \"system\" and \"base\" limits",
				"You need to do little update after. See examples",
			},
			[]string{
				"dozer111",
				"dondo",
				"goro",
			},
			[]Example{
				{
					"auth-sv",
					"dozer111",
					[]string{
						"https://your_git.com/auth-sv/pull-requests/91/overview",
					},
				},
				{
					"payment-sv",
					"goro",
					[]string{
						"https://your_git.com/payment-sv/commits/69a1bb1f09cbe5796f95edf7066be46effcd5ffe",
						"https://your_git.com/payment-sv/commits/69a1bb1f09cbe5796f95edf7066be46effddd12q",
					},
				},
			},
		},
		CurrentVersion: "4.0",
	}

	// ACT

	printer.libraries = []*bumpDependencyConfig{bumpConfig}
	output := printer.Print()

	// ASSERT
	// output show all the fields
	expectedOutput := `(warn)some-private-git/limit-lb
(yellow)	Current: 4.0
(green)	Bump to: 5.0

(yellow)The internal code was splitted to "system" and "base" limits
(yellow)You need to do little update after. See examples

(yellow)	People, who can help u:
(yellow)	  - dozer111
(yellow)	  - dondo
(yellow)	  - goro
(yellow)	Examples:
(yellow)	  auth-sv: (dozer111)
(yellow)	    - https://your_git.com/auth-sv/pull-requests/91/overview
(yellow)	  payment-sv: (goro)
(yellow)	    - https://your_git.com/payment-sv/commits/69a1bb1f09cbe5796f95edf7066be46effcd5ffe
(yellow)	    - https://your_git.com/payment-sv/commits/69a1bb1f09cbe5796f95edf7066be46effddd12q`
	assert.Equal(t, expectedOutput, strings.Join(output, "\n"))
}
