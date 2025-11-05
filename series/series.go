package series

import (
	"strconv"
	"time"

	visualfn "github.com/visual-pivert/go-starter/fn"
)

type SeriesType string

const (
	IntType    SeriesType = "int"
	FloatType             = "float"
	StringType            = "string"
	BoolType              = "bool"
	TimeType              = "time"
)

type Series struct {
	name  string
	data  []any
	stype SeriesType
	len   int
}

func ConvertData(data []any, t SeriesType) []any {
	dateFormats := []string{
		time.RFC3339,
		"2006-01-02",
		"02/01/2006",
		"2006/01/02",
		"02-01-2006",
		"2006-01-02 15:04:05",
	}
	convertedData := make([]any, len(data))
	for i, v := range data {
		switch v.(type) {
		case int, int8, int16, int32, int64:
			if t == FloatType {
				convertedData[i] = float64(v.(int))
			} else {
				convertedData[i] = v
			}
			continue
		case float32, float64:
			convertedData[i] = v
			continue
		case bool:
			convertedData[i] = v
			continue
		case time.Time:
			convertedData[i] = v
			continue
		case string:
			switch t {
			case IntType:
				convertedData[i], _ = strconv.Atoi(v.(string))
				continue
			case FloatType:
				convertedData[i], _ = strconv.ParseFloat(v.(string), 64)
				continue
			case BoolType:
				convertedData[i], _ = strconv.ParseBool(v.(string))
				continue
			case TimeType:
				for _, value := range dateFormats {
					if _, err := time.Parse(value, v.(string)); err == nil {
						convertedData[i] = v.(string)
						break
					}
				}
				continue
			case StringType:
				convertedData[i] = v.(string)
				continue
			}
		}
	}
	return convertedData
}

func newSeries(name string, data []any, t SeriesType) *Series {
	return &Series{
		name:  name,
		data:  data,
		stype: t,
		len:   len(data),
	}
}

func New(name string, data []any, t SeriesType) *Series {
	return newSeries(name, ConvertData(data, t), t)
}

func (s *Series) Copy() *Series {
	return newSeries(s.name, s.data, s.stype)
}

func (s *Series) Len() int {
	return s.len
}

func (s *Series) Name() string {
	return s.name
}

func (s *Series) Rename(name string) *Series {
	return newSeries(name, s.data, s.stype)
}

func (s *Series) Type() SeriesType {
	return s.stype
}

func (s *Series) Get(i int) any {
	return s.data[i]
}

func (s *Series) GetSlice() []any {
	return s.data
}

func (s *Series) Set(i int, data any) *Series {
	d := s.Copy()
	d.data[i] = data
	return d
}

func (s *Series) Append(data any) *Series {
	d := s.Copy()
	d.data = append(s.data, data)
	d.len++
	return d
}

func (s *Series) AppendSlice(data []any) *Series {
	d := s.Copy()
	d.data = append(s.data, data...)
	d.len += len(data)
	return d
}

func (s *Series) FilerToBoolStatement(fn func(any) bool) []bool {
	return visualfn.FilterToBoolStatement(s.data, fn)
}

func (s *Series) ApplyWithBoolStatement(boolSlice []bool, fn func(any) any) *Series {
	d := s.Copy()
	for i, value := range d.data {
		if boolSlice[i] {
			d.data[i] = fn(value)
		}
	}
	return d
}

func (s *Series) IntersectWithBoolStatement(boolSlice []bool) *Series {
	d := newSeries(s.name, []any{}, s.stype)
	for i, value := range s.data {
		if boolSlice[i] {
			d.data = append(d.data, value)
		}
	}
	d.len = len(d.data)
	return d
}
