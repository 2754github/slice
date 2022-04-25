package slice

func Contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if element == v {
			return true
		}
	}

	return false
}
