package stringreverse

func reverString(s []byte) {
	start := 0
	end := len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

/* reverse string
* input string: "i love eth"
* output string: "eth love i" */
func ReverseString(s []byte) {
	reverString(s)
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			reverString(s[start:i])
			start = i + 1
		}
	}
	reverString(s[start:])
}
