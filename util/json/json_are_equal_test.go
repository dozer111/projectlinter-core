package utilJSON_test

import (
	utilJSON "github.com/dozer111/projectlinter-core/util/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONsAreEqual(t *testing.T) {
	someJson := `
{
  "post-install-cmd": [
    "@auto-scripts"
  ],
  "post-update-cmd": [
    "@auto-scripts"
  ],
  "auto-scripts": {
    "cache:clear": "symfony-cmd"
  },
  "php-cs-fixer": "php-cs-fixer fix",
  "rector": "rector process"
}
`
	t.Run("equal", func(t *testing.T) {
		sameJSONWithDifferentOrder := `
{
  "post-install-cmd": [
    "@auto-scripts"
  ],
  "php-cs-fixer": "php-cs-fixer fix",
  "auto-scripts": {
    "cache:clear": "symfony-cmd"
  },
  "rector": "rector process",
  "post-update-cmd": [
    "@auto-scripts"
  ]
}
`

		assert.True(t, utilJSON.JSONsAreEqual([]byte(someJson), []byte(sameJSONWithDifferentOrder)))
	})

	t.Run("different", func(t *testing.T) {
		differentJSON := `
{
  "post-install-cmd": [
    "@auto-scripts"
  ],
  "xvost": "bkb",
  "post-update-cmd": [
    "@auto-scripts"
  ],
  "auto-scripts": {
    "cache:clear": "symfony-cmd"
  },
  "php-cs-fixer": "php-cs-fixer fix",
  "rector": "rector process"
}
`

		assert.False(t, utilJSON.JSONsAreEqual([]byte(someJson), []byte(differentJSON)))
	})
}
