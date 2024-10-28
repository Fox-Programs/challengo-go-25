package piscine

func FindNextPrime(nb int) int {
	result := nb
	i := 2
	if nb <= 2 {
		result = 2
	} else {
		for ; i < nb; i++ {
			if nb%i == 0 {
				// pas nb premier -> indiquer nb premier supérieur à ce nb
				nb++
				i = 2
			}
		}
		result = nb
	}
	return (result)
}
