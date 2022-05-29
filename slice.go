package su

// Includes checks if a slice contains a target item.
func Includes[T comparable](slice []T, target T) (found bool) {
	for _, v := range slice {
		if v == target {
			found = true
		}
	}
	return
}

// Map applies a function to each item in a slice and returns a new slice with the results.
func Map[T, U any](origSlice []T, mapFn func(item T) U) (newSlice []U) {
	for _, v := range origSlice {
		newSlice = append(newSlice, mapFn(v))
	}
	return
}

// IndexOf returns the index of a target item in a slice.
// Returns -1 if the target item is not found.
func IndexOf[T comparable](slice []T, target T) (index int) {
	index = -1
	for i, v := range slice {
		if v == target {
			index = i
			break
		}
	}
	return
}
