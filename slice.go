package su

func Includes[T comparable](slice []T, target T) (found bool) {
	for _, v := range slice {
		if v == target {
			found = true
		}
	}
	return
}

func Map[T, U any](origSlice []T, mapFn func(item T) U) (newSlice []U) {
	for _, v := range origSlice {
		newSlice = append(newSlice, mapFn(v))
	}
	return
}

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
