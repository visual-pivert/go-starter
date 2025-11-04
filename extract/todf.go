package extract

import (
	"reflect"
	"time"

	"github.com/visual-pivert/go-starter/df"
	"github.com/visual-pivert/go-starter/series"
)

func ToDf(raw [][]any, headerIdx int) *df.Df {
	var headers []string
	for i := 0; i < len(raw); i++ {
		headers = append(headers, raw[i][headerIdx].(string))
	}
	out := df.NewDf()
	for i := 0; i < len(headers); i++ {
		var arr []any
		for j := headerIdx + 2; j < len(raw); j++ {
			arr = append(arr, raw[j][i])
		}
		out = out.AddSeries(series.NewSeries(headers[i], arr, GetSliceType(arr)))
	}
	return out
}

func GetSliceType(slice []any) series.SeriesType {
	for _, v := range slice {
		if v == nil {
			continue
		}

		switch val := v.(type) {
		case int, int8, int16, int32, int64:
			return series.IntType
		case float32, float64:
			return series.FloatType
		case bool:
			return series.BoolType
		case string:
			return series.StringType
		case time.Time:
			return series.TimeType
		default:
			// handle numeric values stored as json.Number etc.
			rv := reflect.ValueOf(val)
			if rv.Kind() == reflect.Int || rv.Kind() == reflect.Float64 {
				return series.FloatType
			}
		}
	}

	return series.StringType // default fallback if all nil or unknown
}
