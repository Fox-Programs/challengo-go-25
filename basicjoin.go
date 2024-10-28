package piscine

func BasicJoin(elems []string) string {
	answer := ""
	i := 0
	for range elems {
		answer += elems[i]
		i++
	}
	return answer
}
