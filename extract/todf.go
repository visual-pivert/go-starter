package extract

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/visual-pivert/go-starter/df"
	"github.com/visual-pivert/go-starter/series"
)

func ToDf(raw [][]string, headerIdx int) *df.Df {
	headers := raw[headerIdx]
	out := df.New()
	for i := 0; i < len(headers); i++ {
		var arr []any
		for j := headerIdx + 1; j < len(raw); j++ {
			arr = append(arr, raw[j][i])
		}
		out = out.AddSeries(series.New(headers[i], arr, GetSliceType(arr)))
	}
	return out
}

func GetSliceType(slice []any) series.Type {
	dateFormats := []string{
		time.RFC3339,
		"2006-01-02",
		"02/01/2006",
		"2006/01/02",
		"02-01-2006",
		"2006-01-02 15:04:05",
	}

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

		case time.Time:
			return series.TimeType

		case string:
			s := strings.TrimSpace(val)
			if s == "" {
				continue
			}

			// Try to detect numeric or logical values
			if _, err := strconv.ParseInt(s, 10, 64); err == nil {
				return series.IntType
			}
			if _, err := strconv.ParseFloat(s, 64); err == nil {
				return series.FloatType
			}
			if _, err := strconv.ParseBool(s); err == nil {
				return series.BoolType
			}

			// Try multiple date formats
			for _, format := range dateFormats {
				if _, err := time.Parse(format, s); err == nil {
					return series.TimeType
				}
			}

			return series.StringType

		default:
			rv := reflect.ValueOf(val)
			switch rv.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return series.IntType
			case reflect.Float32, reflect.Float64:
				return series.FloatType
			case reflect.Bool:
				return series.BoolType
			}
		}
	}

	return series.StringType // fallback if all nil or unknown
}
