package su_string

func Substring(s string, start int, args ...int) string {
	if len(args) > 1 {
		panic("Only one extra parameter should be passed")
	}
	runes := []rune(s)
	end := len(runes)

	if len(args) != 0 {
		end = args[0]
	}
	return string(runes[start:end])
}

func CharCodeAt(s string, idx int) int {
	runes := []rune(s)
	return int(runes[idx])
}
