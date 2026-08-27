package piscine

func ReverseMenuIndex(menu []string) []string {
	rev := make([]string, len(menu))
	revind := 0
	for index := len(menu) - 1; index >= 0; index-- {
		rev[revind] = menu[index]
		revind++
	}
	return rev
}
