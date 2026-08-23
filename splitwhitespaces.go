package piscine

func SplitWhiteSpaces(s string) []string {
	words := []string{}
	starting := 0

	for index, value := range s {
		if value == ' ' || value == '\n' || value == '\t' {
			if starting < index {
				words = append(words, s[starting:index])
			}
			starting = index + 1
		}
	}

	if starting < len(s) {
		words = append(words, s[starting:])
	}
	return words
}
