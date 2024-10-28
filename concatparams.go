package piscine

func ConcatParams(args []string) string {
	answer := ""
	for i := 0; i < len(args); i++ {
		answer += (args[i] + "\n")
	}
	return answer[:len(answer)-1] // retire le dernier caractère (passage à la ligne en trop)
}
