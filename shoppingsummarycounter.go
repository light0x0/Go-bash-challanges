package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	card := make(map[string]int)
	word := ""
	for _, value := range str {
		if value == ' ' {
			card[word]++
			word = ""

		} else {
			word += string(value)
		}
	}
	card[word]++
	return card
}
