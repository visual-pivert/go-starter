package series

import (
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

func NewSeries(name string, data []any, t SeriesType) *Series {
	return &Series{
		name:  name,
		data:  data,
		stype: t,
		len:   len(data),
	}
}

func (s *Series) Copy() *Series {
	return NewSeries(s.name, s.data, s.stype)
}

func (s *Series) Len() int {
	return s.len
}

func (s *Series) Name() string {
	return s.name
}

func (s *Series) Rename(name string) *Series {
	return NewSeries(name, s.data, s.stype)
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
	d := NewSeries(s.name, []any{}, s.stype)
	for i, value := range s.data {
		if boolSlice[i] {
			d.data = append(d.data, value)
		}
	}
	d.len = len(d.data)
	return d
}
