package is

func Truthy(value any) bool {
	return !Falsy(value)
}
