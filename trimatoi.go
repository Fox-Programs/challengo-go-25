package piscine

func TrimAtoi(s string) int {
	answer := 0
	isnegative := false
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "-" && answer == 0 {
			isnegative = true
		}
		if rune(s[i]) >= 48 && rune(s[i]) <= 57 {
			answer *= 10
			answer += (int(s[i]) - 48)
		}
	}
	if isnegative {
		answer *= -1
	}
	return answer
}
