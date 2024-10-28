package piscine

func AppendRange(min, max int) []int {
	tab := []int{}
	nb_in_tab := max - min
	for i := 0; i < nb_in_tab; i++ {
		tab = append(tab, min+i)
	}
	if min >= max {
		return []int(nil)
	}
	return tab
}
