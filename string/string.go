package su_string

func Substring(s string, start int, args ...int) string {
	if len(args) > 1 {
		panic("Only one extra parameter should be passed")
	}
	runes := []rune(s)
	end := args[0]
	return string(runes[start:end])
}
