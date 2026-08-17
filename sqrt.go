package piscine

func Sqrt(nb int) int {
	index := 0
	for (index * index) <= nb {
		if index*index == nb {
			return index
		}
		index++
	}

	return 0
}
