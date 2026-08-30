package piscine

func f(a, b int) int {
	return a - b
}

func IsSorted(f func(a, b int) int, a []int) bool {
	big := true
	small := true
	for index := 0; index < len(a)-1; index++ {
		if f(a[index], a[index+1]) > 0 {
			small = false
		}
	}

	for index := 0; index+1 < len(a); index++ {
		if f(a[index], a[index+1]) < 0 {
			big = false
		}
	}

	return small || big
}
