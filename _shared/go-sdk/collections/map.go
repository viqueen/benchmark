package collections

func Map[T, R interface{}](items []T, mapper func(T) R) []R {
	var result []R
	for _, item := range items {
		result = append(result, mapper(item))
	}
	return result
}
