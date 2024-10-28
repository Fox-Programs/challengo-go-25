package piscine

func IsLower(s string) bool {
	board := []rune(s)
	answer := true
	for i := 0; i < len(s); i++ {
		if board[i] < 'a' || board[i] > 'z' {
			answer = false
		}
	}
	return answer
}
