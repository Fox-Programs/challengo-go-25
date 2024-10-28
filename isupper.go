package piscine

func IsUpper(s string) bool {
	board := []rune(s)
	answer := true
	for i := 0; i < len(s); i++ {
		if board[i] <= 'A' || board[i] >= 'Z' {
			answer = false
		}
	}
	return answer
}
