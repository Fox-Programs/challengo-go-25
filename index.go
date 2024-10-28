package piscine

func Index(s string, toFind string) int {
	board_word_1 := []rune(s)
	board_word_2 := []rune(toFind)
	for i := 0; i <= len(board_word_1)-len(board_word_2); i++ {
		if toFind == s[i:i+len(board_word_2)] {
			return i
		}
	}
	return -1
}
