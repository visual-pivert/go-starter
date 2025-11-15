package df

import (
	"testing"

	"github.com/visual-pivert/go-starter/is"
)

func TestDf_FromRaw(t *testing.T) {
	testCases := []struct {
		name           string
		raw            [][]string
		headerId       int
		types          []string
		expectedHeader []string
		expectedShape  []int
	}{
		{
			name: "From raw",
			raw: [][]string{
				{"number", "string"},
				{"1", "a"},
				{"2", "b"},
				{"3", "c"},
			},
			types:          []string{"number", "string"},
			expectedHeader: []string{"number", "string"},
			expectedShape:  []int{3, 2},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(tt *testing.T) {
			got := FromRaw(testCase.raw, testCase.types, testCase.headerId)
			got.Debug()
			if is.SameSlice(testCase.expectedShape, got.Shape()) == false {
				tt.Errorf("Expected %v, got %v", testCase.expectedShape, got.Shape())
			}
			if is.SameSlice(testCase.expectedHeader, got.GetHeaders()) == false {
				tt.Errorf("Expected %v, got %v", testCase.expectedHeader, got.GetHeaders())
			}
		})
	}

}
