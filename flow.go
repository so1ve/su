package su

func If[T any](cond bool, first, last T) T {
	if cond {
		return first
	}
	return last
}
