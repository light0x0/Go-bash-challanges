package piscine

func SortIntegerTable(table []int) {
	for index := 0; index < len(table); index++ {
		for index2 := 0; index2 < len(table)-1; index2++ {
			if table[index2] > table[index2+1] {
				table[index2], table[index2+1] = table[index2+1], table[index2]
			}
		}
	}
}
