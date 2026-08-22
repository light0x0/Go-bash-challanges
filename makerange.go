package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	nums := make([]int, max-min)

	for index := min; index < max ; index++ {
		nums[index - min] = index
	}
	return nums
}
