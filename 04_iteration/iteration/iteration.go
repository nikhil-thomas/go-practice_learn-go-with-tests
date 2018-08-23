package iteration

// Repeat repeats a string 5 times
func Repeat(s string, n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str += s
	}
	return str
}
