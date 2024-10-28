package piscine

func StrRev(s string) string {
	chaine := ""
	for i := len(s); i >= 1; i-- {
		chaine = chaine + string(s[i-1])
	}
	return chaine
}
