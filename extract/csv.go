package extract

import (
	"os"
	"strings"

	"github.com/visual-pivert/go-starter/df"
	"github.com/visual-pivert/go-starter/fn"
)

func parseCsv(content string, sep string) [][]string {
	split := strings.Split(content, "\n")
	split = fn.FilterTruthy(split)
	var out [][]string
	for _, s := range split {
		out = append(out, strings.Split(s, sep))
	}
	return out
}

func Csv(path string, sep string, headerIdx int) *df.Df {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fileStr := string(fileContent)
	parsed := parseCsv(fileStr, sep)
	return df.FromRaw(parsed, headerIdx)
}
