package extract

import (
	"strings"
)

func Csv(content string, sep string) [][]string {
	split := strings.Split(content, "\n")
	var out [][]string
	for _, s := range split {
		out = append(out, strings.Split(s, sep))
	}
	return out
}
