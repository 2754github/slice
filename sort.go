package slice

import (
	"sort"
)

func Sort[T comparable](slice []T, less func(i, j int) bool, stable bool) {
	if stable {
		sort.SliceStable(slice, less)
	} else {
		sort.Slice(slice, less)
	}
}
