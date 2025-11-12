package series

import (
	"fmt"

	fnVisual "github.com/visual-pivert/go-starter/fn"
)

type Series[T comparable] struct {
	data []T
}

func New[T comparable](data []T) Series[T] {
	return Series[T]{data}
}

func (s Series[T]) Append(values []T) Series[T] {
	out := make([]T, 0, len(s.data)+len(values))
	out = append(out, s.data...)
	out = append(out, values...)
	return Series[T]{out}
}

func (s Series[T]) AppendTo(pos int, values []T) Series[T] {
	out := make([]T, 0, len(s.data)+len(values))
	out = append(out, s.data[:pos]...)
	out = append(out, values...)
	out = append(out, s.data[pos:]...)
	return Series[T]{out}
}

func (s Series[T]) Pop() (Series[T], T) {
	out := make([]T, 0, len(s.data)-1)
	out = append(out, s.data[:len(s.data)-1]...)
	return Series[T]{out}, s.data[len(s.data)-1]
}

func (s Series[T]) Shift() (Series[T], T) {
	out := make([]T, 0, len(s.data)-1)
	out = append(out, s.data[1:]...)
	return Series[T]{out}, s.data[0]
}

func (s Series[T]) Remove(index int) Series[T] {
	out := make([]T, 0, len(s.data)-1)
	out = append(out, s.data[:index]...)
	out = append(out, s.data[index+1:]...)
	return Series[T]{out}
}

func (s Series[T]) Range(start int, nbr int) Series[T] {
	out := make([]T, 0, nbr)
	out = append(out, s.data[start:start+nbr]...)
	return Series[T]{out}
}

func (s Series[T]) Len() int {
	return len(s.data)
}

func (s Series[T]) Count() int {
	return len(s.data)
}

func (s Series[T]) Debug() {
	fmt.Println(s.data)
}

func (s Series[T]) ToSlice() []T {
	var out []T
	out = append(out, s.data...)
	return out
}

func (s Series[T]) Filter(fn func(value T) bool) Series[T] {
	out := fnVisual.Filter(s.data, fn)
	return Series[T]{out}
}

func (s Series[T]) FilterI(fn func(value T) bool) Series[int] {
	out := fnVisual.FilterI(s.data, fn)
	return Series[int]{out}
}

func (s Series[T]) Reduce(initialValue T, fn func(last T, curr T, currIndex int) T) Series[T] {
	out := fnVisual.Reduce(s.data, initialValue, fn)
	return Series[T]{out}
}

func (s Series[T]) Map(fn func(value T, index int) T) Series[T] {
	out := fnVisual.Map(s.data, fn)
	return Series[T]{out}
}

func (s Series[T]) MapToBool(fn func(value T, index int) bool) Series[bool] {
	out := fnVisual.Map(s.data, fn)
	return Series[bool]{out}
}

func (s Series[T]) ApplyBoolStatement(boolStatement Series[bool]) Series[T] {
	var out Series[T]
	for i, value := range s.data {
		if boolStatement.GetValue(i) {
			out = out.Append([]T{value})
		}
	}
	return out
}

func (s Series[T]) ApplyOrderStatement(orderStatement Series[int]) Series[T] {
	var out Series[T]
	for _, index := range orderStatement.ToSlice() {
		out = out.Append([]T{s.data[index]})
	}
	return out
}

func (s Series[T]) CountValue(value T) int {
	counter := 0
	for _, v := range s.data {
		if v == value {
			counter++
		}
	}
	return counter
}

func (s Series[T]) GetValue(index int) T {
	return s.data[index]
}

func (s Series[T]) SetValue(index int, value T) Series[T] {
	s.data[index] = value
	return Series[T]{s.data}
}

func (s Series[T]) Reverse() Series[T] {
	out := fnVisual.Reverse(s.data)
	return Series[T]{out}
}

func (s Series[T]) Agg(initialValue T, fn func(last T, curr T, currIndex int) T) T {
	out := fnVisual.Reduce(s.data, initialValue, fn)
	return out[len(out)-1]
}

func (s Series[T]) Any(fn func(value T) bool) bool {
	out := fnVisual.Any(s.data, fn)
	return out
}

func (s Series[T]) All(fn func(value T) bool) bool {
	out := fnVisual.All(s.data, fn)
	return out
}

func (s Series[T]) IndexOf(value T) int {
	out := fnVisual.IndexOf(value, s.data)
	return out
}
