package piscine

func Sqrt(nb int) int {
	result := 0
	if nb <= 0 {
		result = 0
	} else {
		for i := 1; i <= nb; i++ {
			if float32(nb)/float32(i) == float32(i) {
				result = i
			}
		}
	}
	return result
}
