package piscine

func BasicAtoi2(s string) int {
	t := 0

	for i := 0; i < len(s); i++ {
		if int(s[i]-'0') >= 0 && int(s[i]-'0') <= 9 {
			t = t * 10
			t += int(s[i] - '0')
		} else {
			return 0
		}
	}
	return t
}
