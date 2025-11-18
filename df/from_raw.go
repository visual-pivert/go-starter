package df

import "github.com/visual-pivert/go-starter/series"

// FromRaw creates a dataframe from raw data.
// Example:
//
//	df.FromRaw([][]string{
//		{"a", "b", "c"},
//		{"1", "2", "3"},
//		{"4", "5", "6"},
//	}, nil, 0)
//	df.Debug()
//
// Output:
// | a | b | c |
// ---
// | 1 | 2 | 3 |
// | 4 | 5 | 6 |
func FromRaw(data [][]string, types []string, headerId int) *Dataframe {
	if len(data) == 0 || headerId < 0 || headerId >= len(data) {
		return New(nil, []string{})
	}

	headers := data[headerId]
	cols := len(headers)

	startRow := headerId + 1
	if startRow > len(data) {
		return New(nil, headers)
	}
	rows := len(data) - startRow

	dataframe := make([][]any, cols)
	for c := 0; c < cols; c++ {
		col := make([]any, rows)
		for r := 0; r < rows; r++ {
			row := data[startRow+r]
			if c < len(row) {
				col[r] = row[c]
			} else {
				col[r] = ""
			}
		}
		dataframe[c] = col
	}

	if len(types) != cols {
		types = make([]string, cols)
		for i := 0; i < cols; i++ {
			types[i] = "string"
		}
	}

	newDf := New(nil, []string{})
	for idx, col := range dataframe {
		newDf.Append(series.New[any](col, types[idx]), headers[idx])
	}
	return newDf
}
