package collections

func MapValues[V interface{}](m map[string]V) []V {
	var result []V
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
