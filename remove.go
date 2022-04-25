package slice

func Remove[T comparable](slice []T, element T) []T {
	removed := make([]T, 0, len(slice))
	for _, v := range slice {
		if element == v {
			continue
		}
		removed = append(removed, v)
	}

	return removed
}
