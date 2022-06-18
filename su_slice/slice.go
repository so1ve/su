package su_slice

import "reflect"

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

// Includes checks if a slice contains a target item.
func Includes[T comparable](slice []T, target T) bool {
	return IndexOf(slice, target) != -1
}

// IndexOfReflect is similar to IndexOf but uses runtime reflection.
func IndexOfReflect[T any](slice []T, target T) (index int) {
	index = -1
	for i, v := range slice {
		if reflect.DeepEqual(v, target) {
			index = i
			break
		}
	}
	return
}

// IncludesReflect is similar to Includes but uses runtime reflection.
func IncludesReflect[T any](slice []T, target T) bool {
	return IndexOfReflect(slice, target) != -1
}

func Every[T comparable](slice []T, everyFn func(item T, idx int) bool) bool {
	for i, v := range slice {
		if !everyFn(v, i) {
			return false
		}
	}
	return true
}

func Some[T comparable](slice []T, someFn func(item T, idx int) bool) bool {
	for i, v := range slice {
		if someFn(v, i) {
			return true
		}
	}
	return false
}
