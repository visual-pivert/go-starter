package series

// series package contains a Series and many utilities for working with.
// Series is a supercharged version of slices.

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	fnVisual "github.com/visual-pivert/go-starter/fn"
	"github.com/visual-pivert/go-starter/is"
)

type Series[T any] struct {
	data []T
	t    string // "string" or "number" or "date" or "float" or "bool"
}

// New creates a new Series.
// Param t must be one of "string", "number", "date", "float" or "bool".
// Examples:
//
//	series.New([]int{1, 2, 3}, "number") // return Series of type number
func New[T any](data []T, t string) Series[T] {
	typePossibilities := []string{"string", "number", "date", "float", "bool"}
	if is.In(t, typePossibilities) == false {
		panic("type not supported")
	}
	coerced, ok := coerceIfAnySlice[T](data, t)
	if ok {
		return Series[T]{coerced, t}
	}
	return Series[T]{data, t}
}

// coerceIfAnySlice attempts to coerce a slice of interface values into a slice of a specified type.
// If coercion succeeds, it returns the coerced slice and true; otherwise, it returns nil and false.
func coerceIfAnySlice[T any](data []T, t string) ([]T, bool) {
	rv := reflect.ValueOf(data)
	if rv.Kind() != reflect.Slice {
		return nil, false
	}
	elemT := rv.Type().Elem()
	if elemT.Kind() != reflect.Interface {
		return nil, false
	}
	dstElem := reflect.TypeOf((*T)(nil)).Elem()
	if dstElem.Kind() != reflect.Interface {
		return nil, false
	}
	n := rv.Len()
	out := reflect.MakeSlice(reflect.SliceOf(dstElem), n, n)
	for i := 0; i < n; i++ {
		elem := rv.Index(i)
		if elem.Kind() == reflect.Interface && elem.IsNil() {
			out.Index(i).Set(reflect.ValueOf(zeroForType(t)))
			continue
		}
		v := elem.Interface() // dynamic value
		if s, ok := v.(string); ok && t != "string" {
			if conv, okc := convertStringToType(strings.TrimSpace(s), t); okc {
				out.Index(i).Set(reflect.ValueOf(conv))
				continue
			}
			out.Index(i).Set(reflect.ValueOf(zeroForType(t)))
			continue
		}
		out.Index(i).Set(reflect.ValueOf(v))
	}
	return out.Interface().([]T), true
}

// convertStringToType converts a string `s` into a specified type `t` ("number", "float", "bool", "date", or default string).
// Returns the converted value and a boolean indicating success or failure of the conversion.
func convertStringToType(s string, t string) (any, bool) {
	switch t {
	case "number":
		if s == "" { // empty -> zero
			return 0, true
		}
		if i, err := strconv.Atoi(s); err == nil {
			return i, true
		}
		return nil, false
	case "float":
		if s == "" {
			return 0.0, true
		}
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			return f, true
		}
		return nil, false
	case "bool":
		if s == "" {
			return false, true
		}
		if b, err := strconv.ParseBool(s); err == nil {
			return b, true
		}
		return nil, false
	case "date":
		return s, true
	default:
		return s, true
	}
}

// zeroForType returns the zero value for the specified type as a string ("number", "float", "bool", "date", or default string).
func zeroForType(t string) any {
	switch t {
	case "number":
		return 0
	case "float":
		return 0.0
	case "bool":
		return false
	case "date":
		return ""
	default:
		return ""
	}
}

// Append adds values to the end of the Series.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s = s.Append([]int{4, 5}) // return Series of type number with values [1, 2, 3, 4, 5]
//	s.Debug() // [1, 2, 3, 4, 5]
func (s Series[T]) Append(values []T) Series[T] {
	out := make([]T, 0, len(s.data)+len(values))
	out = append(out, s.data...)
	out = append(out, values...)
	return Series[T]{out, s.t}
}

// AppendTo inserts values at a given position in the Series.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s = s.AppendTo(1, []int{4, 5}) // return Series of type number with values [1, 4, 5, 2, 3]
//	s.Debug() // [1, 4, 5, 2, 3]
func (s Series[T]) AppendTo(pos int, values []T) Series[T] {
	out := make([]T, 0, len(s.data)+len(values))
	out = append(out, s.data[:pos]...)
	out = append(out, values...)
	out = append(out, s.data[pos:]...)
	return Series[T]{out, s.t}
}

// Pop removes the last value from the Series and returns it.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s, value := s.Pop() // return Series of type number with values [1, 2] and value 3
//	s.Debug() // [1,2]
func (s Series[T]) Pop() (Series[T], T) {
	out := make([]T, 0, len(s.data)-1)
	out = append(out, s.data[:len(s.data)-1]...)
	return Series[T]{out, s.t}, s.data[len(s.data)-1]
}

// Shift removes the first value from the Series and returns it.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s, value := s.Shift() // return Series of type number with values [2, 3] and value 1
//	s.Debug() // [2,3]
func (s Series[T]) Shift() (Series[T], T) {
	out := make([]T, 0, len(s.data)-1)
	out = append(out, s.data[1:]...)
	return Series[T]{out, s.t}, s.data[0]
}

// Remove removes a value at a given position in the Series.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s = s.Remove(1) // return Series of type number with values [1, 3]
func (s Series[T]) Remove(index int) Series[T] {
	out := make([]T, 0, len(s.data)-1)
	out = append(out, s.data[:index]...)
	out = append(out, s.data[index+1:]...)
	return Series[T]{out, s.t}
}

// Range returns a new Series with values from a given range.
// Examples:
//
//	s := series.New([]int{1, 2, 3, 4, 5}, "number")
//	s = s.Range(1, 3) // return Series of type number with values [2, 3] (from index 1 and take 3 values from there)
func (s Series[T]) Range(start int, nbr int) Series[T] {
	out := make([]T, 0, nbr)
	out = append(out, s.data[start:start+nbr]...)
	return Series[T]{out, s.t}
}

// Len returns the length of the Series.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s.Len() // return 3
func (s Series[T]) Len() int {
	return len(s.data)
}

// Count returns the number of elements in the Series.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s.Count() // return 3 (same as s.Len())
func (s Series[T]) Count() int {
	return len(s.data)
}

// Debug prints the Series to the console.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s.Debug() // [1, 2, 3]
func (s Series[T]) Debug() {
	parts := make([]string, len(s.data))
	for i, v := range s.data {
		parts[i] = fmt.Sprintf("%v", v)
	}
	fmt.Printf("[%s]\n", strings.Join(parts, ", "))
}

// Type returns the type of the Series.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s.Type() // return "number"
func (s Series[T]) Type() string {
	return s.t
}

// ToSlice returns a slice of the Series.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s.ToSlice() // return []int{1, 2, 3}
func (s Series[T]) ToSlice() []T {
	var out []T
	out = append(out, s.data...)
	return out
}

// Filter returns a new Series containing elements that satisfy the given filtering function.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s = s.Filter(func(value int) bool { return value > 1 }) // return Series of type number with values [2, 3]
//	s.Debug() // [2, 3]
func (s Series[T]) Filter(fn func(value T) bool) Series[T] {
	out := fnVisual.Filter(s.data, fn)
	return Series[T]{out, s.t}
}

// FilterI returns a new Series containing indices of elements that satisfy the given filtering function.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s = s.FilterI(func(value int) bool { return value > 1 }) // return Series of type number with values [1, 2] (indices)
func (s Series[T]) FilterI(fn func(value T) bool) Series[int] {
	out := fnVisual.FilterI(s.data, fn)
	return Series[int]{out, s.t}
}

// Reduce return a new Series that contains the cumulative result.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s = s.Reduce(0, func(last int, curr int, currIndex int) int { return last + curr }) // return Series of type number with values [1, 3, 6]
//	s.Debug() // [1, 3, 6]
func (s Series[T]) Reduce(initialValue T, fn func(last T, curr T, currIndex int) T) Series[T] {
	out := fnVisual.Reduce(s.data, initialValue, fn)
	return Series[T]{out, s.t}
}

// Map returns a new Series with the results of calling a provided function on every element in the calling Series.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s = s.Map(func(value int) int { return value * 2 }) // return Series of type number with values [2, 4, 6]
func (s Series[T]) Map(fn func(value T, index int) T) Series[T] {
	out := fnVisual.Map(s.data, fn)
	return Series[T]{out, s.t}
}

// MapToBool returns a new bool Series with the results of calling a provided function on every element in the calling Series.
// True if the function returns true, false otherwise.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s = s.MapToBool(func(value int) bool { return value > 1 }) // return Series of type bool with values [false, true, true]
func (s Series[T]) MapToBool(fn func(value T, index int) bool) Series[bool] {
	out := fnVisual.Map(s.data, fn)
	return Series[bool]{out, s.t}
}

// ApplyBoolStatement returns a new Series where true values are kept and false values are removed.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	boolStatement := series.New([]bool{true, false, true}, "bool")
//	s = s.ApplyBoolStatement(boolStatement) // return Series of type number with values [1, 3]
func (s Series[T]) ApplyBoolStatement(boolStatement Series[bool]) Series[T] {
	var out Series[T]
	for i, value := range s.data {
		if boolStatement.GetValue(i) {
			out = out.Append([]T{value})
		}
	}
	return out
}

// ApplyOrderStatement returns a new Series with elements in the order of the given Series.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	orderStatement := series.New([]int{2, 0, 1}, "number")
//	s = s.ApplyOrderStatement(orderStatement) // return Series of type number with values [3, 1, 2]
func (s Series[T]) ApplyOrderStatement(orderStatement Series[int]) Series[T] {
	var out Series[T]
	for _, index := range orderStatement.ToSlice() {
		out = out.Append([]T{s.data[index]})
	}
	return out
}

// CountValue counts the number of occurrences of the specified value in the Series.
// It returns the count as an integer.
// Examples:
//
//	s := series.New([]int{1, 2, 3, 1, 2}, "number")
//	s.CountValue(1) // return 2
func (s Series[T]) CountValue(value T) int {
	counter := 0
	for _, v := range s.data {
		if reflect.DeepEqual(v, value) {
			counter++
		}
	}
	return counter
}

// GetValue retrieves the value at the specified index from the Series.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s.GetValue(1) // return 2
func (s Series[T]) GetValue(index int) T {
	return s.data[index]
}

// SetValue sets the value at the specified index in the Series.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s = s.SetValue(1, 4) // return Series of type number with values [1, 4, 3]
func (s Series[T]) SetValue(index int, value T) Series[T] {
	s.data[index] = value
	return Series[T]{s.data, s.t}
}

// Reverse returns a new Series with the values in reverse order.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s = s.Reverse() // return Series of type number with values [3, 2, 1]
func (s Series[T]) Reverse() Series[T] {
	out := fnVisual.Reverse(s.data)
	return Series[T]{out, s.t}
}

// Agg returns the result of applying the provided aggregation function to all elements in the Series.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s.Agg(0, func(last int, curr int, currIndex int) int { return last + curr }) // return 6
func (s Series[T]) Agg(initialValue T, fn func(last T, curr T, currIndex int) T) T {
	out := fnVisual.Reduce(s.data, initialValue, fn)
	return out[len(out)-1]
}

// Any returns true if the provided function returns true for any element in the Series.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s.Any(func(value int) bool { return value > 1 }) // return true
func (s Series[T]) Any(fn func(value T) bool) bool {
	out := fnVisual.Any(s.data, fn)
	return out
}

// All returns true if the provided function returns true for all elements in the Series.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s.All(func(value int) bool { return value > 0 }) // return true
func (s Series[T]) All(fn func(value T) bool) bool {
	out := fnVisual.All(s.data, fn)
	return out
}

// IndexOf returns the first index at which a given element can be found in the Series, or -1 if it is not present.
// Examples:
//
//	s := series.New([]int{1, 2, 3}, "number")
//	s.IndexOf(2) // return 1
func (s Series[T]) IndexOf(value T) int {
	out := fnVisual.IndexOf(value, s.data)
	return out
}
