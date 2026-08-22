package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	nums := []int{}

	for index := min; index < max; index++ {
		nums = append(nums, index)
	}
	return nums
}
