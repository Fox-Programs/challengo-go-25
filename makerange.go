package piscine

func MakeRange(min, max int) []int {
	nb_in_tab := max - min
	if min >= max {
		return []int(nil)
	}
	tab := make([]int, nb_in_tab)
	for i := 0; i < nb_in_tab; i++ {
		tab[i] = i + min
	}
	return tab
}
