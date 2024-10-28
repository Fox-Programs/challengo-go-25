package piscine

func Capitalize(s string) string {
	answer := ""
	in_a_word := false
	for i := 0; i < len(s); i++ {
		if in_a_word == false && ((s[i] >= '0' && s[i] <= '9') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z')) {
			in_a_word = true
			if s[i] >= 'a' && s[i] <= 'z' {
				answer += string(rune(s[i] - 32))
			} else {
				answer += string(s[i])
			}
		} else {
			if s[i] >= 'A' && s[i] <= 'Z' {
				answer += string(rune(s[i] + 32))
			} else {
				answer += string(s[i])
			}
			if !(s[i] >= '0' && s[i] <= '9') && !(s[i] >= 'A' && s[i] <= 'Z') && !(s[i] >= 'a' && s[i] <= 'z') {
				in_a_word = false
			}
		}
	}
	return answer
}
