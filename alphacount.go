package piscine

func AlphaCount(s string) int {
	count := 0
	for alpha := 'a'; alpha <= 'z'; alpha++ {
		for i := 0; i < len(s); i++ {
			if string(s[i]) == string(alpha) {
				count++
			}
		}
	}
	for Alpha := 'A'; Alpha <= 'Z'; Alpha++ {
		for i := 0; i < len(s); i++ {
			if string(s[i]) == string(Alpha) {
				count++
			}
		}
	}
	return count
}
