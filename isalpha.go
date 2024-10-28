package piscine

func IsAlpha(s string) bool {
	for _, letter := range s {
		if !((letter >= 'a' && letter <= 'z') || (letter >= '0' && letter <= '9') || (letter >= 'A' && letter <= 'Z')) {
			return false
		}
	}
	return true
}
