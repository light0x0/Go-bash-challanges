package piscine

func IterativePower(nb int, power int) int {
	total := 1

	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}

	for index := 1; index < power+1; index++ {
		total = total * nb
	}

	return total
}
