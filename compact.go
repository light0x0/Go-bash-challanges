package piscine

func Compact(ptr *[]string) int {
	rm := []string{}
	for _, value := range *ptr {
		if value != "" {
			rm = append(rm, value)
		}
	}
	*ptr = rm
	return len(rm)
}
