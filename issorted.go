package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for index := 0; index < len(a)-1; index++ {
		if f(a[index], a[index+1]) > 0 {
			return false
		}
	}

	return true
}
