package piscine

func IsPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		if rune(s[i]) < 32 || rune(s[i]) > 128 {
			return false
		}
	}
	return true
}
