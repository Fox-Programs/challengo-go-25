package piscine

func ToLower(s string) string {
	letter := ""
	for i := 0; i < len(s); i++ {
		if !(rune(s[i]) >= 'A' && rune(s[i]) <= 'Z') {
			letter += string(rune(s[i]))
		} else {
			letter += string(rune((s[i]) + 32))
		}
	}
	return letter
}
