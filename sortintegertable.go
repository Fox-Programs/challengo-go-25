package piscine

func SortIntegerTable(table []int) {
	for i := 1; i <= len(table)-1; i++ {
		key := table[i]
		j := i - 1
		for j >= 0 && table[j] > key {
			table[j+1] = table[j]
			j--
		}
		table[j+1] = key
	}
}
