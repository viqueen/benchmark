package collections

func Intersection[T comparable](left, right []T) []T {
	// Create a map to store the elements of the first slice
	elementMap := make(map[T]bool)
	for _, value := range left {
		elementMap[value] = true
	}

	// Create a slice to store the intersection result
	var result []T

	// Check for elements in the second slice that are also in the map
	for _, value := range right {
		if elementMap[value] {
			result = append(result, value)
			// Optionally, remove the element from the map to avoid duplicates
			delete(elementMap, value)
		}
	}

	return result
}

func IntersectionOfAny[T any](left, right []T, comparator func(T, T) bool) []T {
	var result []T

	// Iterate over the first slice and check for each element in the second slice
	for _, v1 := range left {
		for _, v2 := range right {
			if comparator(v1, v2) {
				result = append(result, v1)
				break
			}
		}
	}

	return result
}
