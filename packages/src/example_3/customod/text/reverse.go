package text

// Reverse a text.
// Taken from http://rosettacode.org/wiki/Reverse_a_string#Go
func ReverseText(text string) string {
	r := make([]rune, len(text))
	start := len(text)
	for _, c := range text {
		start--
		r[start] = c
	}
	return string(r[start:])
}
