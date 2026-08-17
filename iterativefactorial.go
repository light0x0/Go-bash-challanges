package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return -1
	}
	if nb == 0 || nb == 1 {
		return 1
	}

	total := 1
	for index := 2; index <= nb; index++ {
		total = total * index
	}

	return total
}
