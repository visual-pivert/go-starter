package series

func Aggregate(s *Series, initialValue any, fn func(last any, curr any, currIndex int) any) any {
	value := initialValue
	for i := 0; i < s.Len(); i++ {
		value = fn(value, s.data[i], i)
	}
	return value
}
