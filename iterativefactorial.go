package piscine

func IterativeFactorial(nb int) int {
	i := 0
	if nb > 0 && nb <= 20 {
		for i = nb - 1; i > 0; i-- {
			nb *= i
		}
	} else if nb == 0 {
		nb = 1
	} else if nb < 0 || nb > 20 {
		nb = 0
	}
	return nb
}
