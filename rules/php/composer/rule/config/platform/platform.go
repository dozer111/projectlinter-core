package platform

import (
	"encoding/json"
	"strings"
)

func platformPrettyJson(platforms map[string]string) []string {
	type config struct {
		Platform map[string]string `json:"platform"`
	}

	type result struct {
		Config config `json:"config"`
	}

	res := result{
		Config: config{
			Platform: platforms,
		},
	}

	b, _ := json.MarshalIndent(res, "", "	")

	return strings.Split(string(b), "\n")
}
