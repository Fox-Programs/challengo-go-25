package piscine

func LastRune(s string) rune {
	i := len(s) - 1
	return []rune(s)[i]
}
