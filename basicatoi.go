package piscine

func BasicAtoi(s string) int {
	t := 0

	for i := 0; i < len(s); i++ {
		t = t * 10
		t += int(s[i] - '0')
	}
	return t
}
