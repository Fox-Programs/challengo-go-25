package piscine

func ToUpper(s string) string {
	letter := ""
	for i := 0; i < len(s); i++ {
		if !(rune(s[i]) >= 'a' && rune(s[i]) <= 'z') {
			letter += string(rune(s[i]))
		} else {
			letter += string(rune((s[i]) - 32))
		}
	}
	return letter
}
