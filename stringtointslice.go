package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}
	slice := []int{}
	for _, value := range str {
		slice = append(slice, int(value))
	}
	return slice
}
