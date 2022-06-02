package su_flow

// If is a ternary conditional operator function
// Returns the "first" argument if the condition is true, otherwise the "second" argument
func If[T any](cond bool, first, last T) T {
	if cond {
		return first
	}
	return last
}
