package fn

import "github.com/visual-pivert/go-starter/is"

// Filter returns a new slice with all elements that satisfy the predicate function.
// Examples:
//
//	fn.Filter([]int{1, 2, 3}, func(v int) bool { return v > 1 }) // [2, 3]

func Filter[T any](slice []T, fn func(T) bool) []T {
	var out []T
	for _, value := range slice {
		if fn(value) {
			out = append(out, value)
		}
	}
	return out
}

// FilterI returns a new slice that contains indices that satisfies the predicate function.
// Examples:
//
//	fn.FilterI([]int{1, 2, 3}, func(v int) bool { return v > 1 }) // [1, 2] (indices)
func FilterI[T any](slice []T, fn func(T) bool) []int {
	var out []int
	for i, value := range slice {
		if fn(value) {
			out = append(out, i)
		}
	}
	return out
}

// FilterTruthy returns a new slice with all truthy values.
// Examples:
//
//	fn.FilterTruthy([]int{1, 0, 2}) // [1, 2]
func FilterTruthy[T any](slice []T) []T {
	var out []T
	for _, value := range slice {
		if !is.Falsy(value) {
			out = append(out, value)
		}
	}
	return out
}

// FilterITruthy returns a new slice that contains indices of truthy values.
// Examples:
//
//	fn.FilterITruthy([]int{1, 0, 2}) // [0, 2] (indexes)
func FilterITruthy[T any](slice []T) []int {
	var out []int
	for i, value := range slice {
		if !is.Falsy(value) {
			out = append(out, i)
		}
	}
	return out
}

// FilterToBoolStatement returns a new slice of bool with the same length as the input slice.
// True if the predicate function returns true, false otherwise.
// Examples:
//
//	fn.FilterToBoolStatement([]int{1, 2, 3}, func(v int) bool { return v > 1 }) // [false, true, true]
func FilterToBoolStatement[T any](slice []T, fn func(T) bool) []bool {
	var out []bool
	for _, value := range slice {
		out = append(out, fn(value))
	}
	return out
}
