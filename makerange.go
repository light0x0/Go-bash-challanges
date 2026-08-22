package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	nums := make([]int, max-min)

	for index := 0; index < len(nums); index++ {
		nums[index] = min + index
	}
	return nums
}
