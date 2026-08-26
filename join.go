package piscine

func Join(strs []string, sep string) string {
	str := ""
	for index, value := range strs {
		if index != len(strs)-1 {
			str += value + sep
		} else {
			str += value + "\n"
		}
	}

	return str
}
