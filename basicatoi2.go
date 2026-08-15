package piscine

func BasicAtoi2(s string) int {
	nums := 0

	for _, value := range s {

		if value < '0' || value > '9' {
			return 0
		}

		nums = nums*10 + int(value-'0')
	}
	return nums
}
