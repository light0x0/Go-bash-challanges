package piscine

func ShoppingListSort(slice []string) []string {
	salsa := []string(slice)

	for index := 0; index < len(salsa); index++ {
		for indexa := index + 1; indexa < len(salsa); indexa++ {
			if len(salsa[index]) > len(salsa[indexa]) {
				salsa[index], salsa[indexa] = salsa[indexa], salsa[index]
			}
		}
	}
	return salsa
}
