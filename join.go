package piscine

func Join(strs []string, sep string) string {
	answer := ""
	i := 0
	for range strs {
		answer += strs[i] + sep
		i++
	}
	answer = answer[:len(answer)-len(sep)]
	return answer
}
