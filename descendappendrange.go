package piscine

func DescendAppendRange(max, min int) []int {
	salsa := []int{}

	if max <= min {
		return salsa
	}

	for index := max; index > min; index-- {
		salsa = append(salsa, index)
	}
	return salsa
}
