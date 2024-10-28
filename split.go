package piscine

func Split(s, sep string) []string {
	size := 0
	for i := range sep {
		size = i + 1
	}
	size2 := 0
	for i := range s {
		size2 = i + 1
	}
	for i := 0; i < size2-size; i++ {
		if s[i:i+size] == sep {
			s = s[:i] + " " + s[i+size:]
			size2 -= size
		}
	}
	return Split_2(s)
}

func Split_2(s string) []string {
	var answer []string
	var mot string
	for _, r := range s {
		if r != ' ' {
			mot += string(r)
		} else {
			if mot != "" {
				answer = append(answer, mot)
				mot = ""
			}
		}
	}
	if mot != "" {
		answer = append(answer, mot)
	}

	return answer
}
