package su_string

func Substring(s string, start, end int) string {
	runes := []rune(s)
	return string(runes[start:end])
}
