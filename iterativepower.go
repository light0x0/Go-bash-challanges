package piscine

func IterativePower(nb int, power int) int {
	total := 1

	if nb >= 21 || power < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}

	for index := 2; index <= power+1; index++ {
		total = total * nb
	}

	return total
}
