package extract

import (
	"testing"

	"github.com/visual-pivert/go-starter/fn"
	"github.com/visual-pivert/go-starter/is"
)

func TestCsv(t *testing.T) {
	testCases := []struct {
		name     string
		value    string
		sep      string
		expected [][]string
	}{
		{"simple csv", "name;weight;age\nJunior;55.5;18\nMino;150.12;25", ";", [][]string{{"name", "weight", "age"}, {"Junior", "55.5", "18"}, {"Mino", "150.12", "25"}}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := ParseCsv(testCase.value, testCase.sep)
			checkingSlice := fn.Map(got, func(t []string, idx int) any {
				return is.SameSlice(t, testCase.expected[idx])
			})
			isSame := fn.Any(checkingSlice, func(a any) bool {
				return a == true
			})
			if !isSame {
				tt.Errorf("Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}
