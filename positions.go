package slice

func Positions[T comparable](slice []T, element T) []int {
	result := make([]int, 0, len(slice))
	for i, v := range slice {
		if element == v {
			result = append(result, i)
		}
	}

	return result
}
