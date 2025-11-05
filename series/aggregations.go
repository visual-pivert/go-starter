package series

func Aggregate[T any](s *Series, initialValue T, fn func(last any, curr any, currIndex int) T) T {
	value := initialValue
	for i := 0; i < s.Len(); i++ {
		value = fn(value, s.data[i], i)
	}
	return value
}
