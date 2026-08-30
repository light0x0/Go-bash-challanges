package piscine

func comp(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	for index := 0; index < len(a)-1; index++ {
		if f(a[index], a[index+1]) < 0 {
			return false
		}
	}

	return true
}
