package extract

import (
	"os"
	"strings"

	"github.com/visual-pivert/go-starter/df"
	"github.com/visual-pivert/go-starter/fn"
)

// parseCsv splits a csv string into rows.
func parseCsv(content string, sep string) [][]string {
	split := strings.Split(content, "\n")
	split = fn.FilterTruthy(split)
	var out [][]string
	for _, s := range split {
		out = append(out, strings.Split(s, sep))
	}
	return out
}

// Csv reads a csv file and returns a dataframe(check out df package for more info).
// Examples:
//
//	extract.Csv("data.csv", ",", 0, []string{"int", "string"}) // return dataframe
func Csv(path string, sep string, headerIdx int, types []string) *df.Dataframe {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fileStr := string(fileContent)
	parsed := parseCsv(fileStr, sep)
	return df.FromRaw(parsed, types, headerIdx)
}
