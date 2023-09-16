package core

func removElementFromSliceeWithIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func isExistsInSlice[T any](list []T, element T, equals func(e1, e2 T) bool) bool {
	for _, e := range list {
		if equals(e, element) {
			return true
		}
	}
	return false
}

func indexOf[T any](list []T, element T, equals func(e1, e2 T) bool) int {
	for i, e := range list {
		if equals(e, element) {
			return i
		}
	}
	return -1
}
