package piscine

func Abort(a, b, c, d, e int) int {
	slice := []int{a, b, c, d, e}
	for index := 0; index < len(slice); index++ {
		for indexa := index + 1; indexa < len(slice); indexa++ {
			if slice[index] > slice[indexa] {
				slice[index], slice[indexa] = slice[indexa], slice[index]
			}
		}
	}

	//mid := len(slice) / 2

	return slice[2]
}
