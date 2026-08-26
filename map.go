package piscine

func Map(f func(int) bool, a []int) []bool {
	check := make([]bool, len(a))
	for index, value := range a {
		check[index] = f(value)
	}
	return check
}
