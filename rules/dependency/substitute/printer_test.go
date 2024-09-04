package substitute

import (
	"strings"
	"testing"

	"github.com/dozer111/projectlinter-core/util/painter"
	"github.com/stretchr/testify/assert"
)

func TestPrinterOnFullCorrectExample(t *testing.T) {
	painter.CurrentPainterState = painter.Fake

	// ARRANGE: we need to substitute some dependencies
	printer := substituteLibraryPrinter{}

	// ACT
	printer.libraries = []Library{
		{
			Name:     "symfony/swiftmailer-bundle",
			ChangeTo: "symfony/mailer",
		},
		{
			Name:     "some-workspace1/library1",
			ChangeTo: "wrkspace/lb2",
			Description: []string{
				"Library1 is corrupted",
				"Change it ASAP!!!",
				"Read https://.....",
			},
			ResponsiblePersons: []string{
				"dozer111",
				"tsukerberg",
				"onyxia",
			},
			Examples: []Example{
				{
					"auth-sv",
					"wager",
					[]string{
						"something",
						"really important to know",
						"about this code",
					},
					[]string{
						"https://your.git.com/auth-sv/commits/69a1bb1f09cbe5796f95edf7066be46effcd5ffe",
						"https://your.git.com/auth-sv/commits/69a1bb1f09cbe5796f922237066be46effcd5ffe",
						"https://your.git.com/auth-sv/commits/69a1bb1f09cbegggg121edf7066be46effcd5ffe",
					},
				},
				{
					"payment-sv",
					"dozer111",
					nil,
					[]string{
						"https://your.git.com/payment-sv/pull-requests/100",
					},
				},
			},
		},
		{
			Name:     "rector/old",
			ChangeTo: "rector/new",
			Description: []string{
				"Rector has released new major version in separate project",
				"See https://rector.org/new-major",
			},
			ResponsiblePersons: []string{
				"dozer111",
				"tokio",
				"maexna",
			},
		},
	}
	output := printer.Print()

	// ASSERT
	// output show all the fields
	expectedOutput := `(warn)Dependency "symfony/swiftmailer-bundle" is deprecated.
(yellow)  Change it to "symfony/mailer"
(yellow)=================================================================================================
(warn)Dependency "some-workspace1/library1" is deprecated.
(yellow)  Change it to "wrkspace/lb2"

(yellow)Library1 is corrupted
(yellow)Change it ASAP!!!
(yellow)Read https://.....

(yellow)People, who can help u:
(yellow)	- dozer111
(yellow)	- tsukerberg
(yellow)	- onyxia
(yellow)Examples:
(yellow)	  auth-sv: (wager)

(yellow)	    something
(yellow)	    really important to know
(yellow)	    about this code

(yellow)	    - https://your.git.com/auth-sv/commits/69a1bb1f09cbe5796f95edf7066be46effcd5ffe
(yellow)	    - https://your.git.com/auth-sv/commits/69a1bb1f09cbe5796f922237066be46effcd5ffe
(yellow)	    - https://your.git.com/auth-sv/commits/69a1bb1f09cbegggg121edf7066be46effcd5ffe
(yellow)	  payment-sv: (dozer111)
(yellow)	    - https://your.git.com/payment-sv/pull-requests/100
(yellow)=================================================================================================
(warn)Dependency "rector/old" is deprecated.
(yellow)  Change it to "rector/new"

(yellow)Rector has released new major version in separate project
(yellow)See https://rector.org/new-major

(yellow)People, who can help u:
(yellow)	- dozer111
(yellow)	- tokio
(yellow)	- maexna`
	assert.Equal(t, expectedOutput, strings.Join(output, "\n"))
}
