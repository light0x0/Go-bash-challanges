package piscine

func ConcatParams(args []string) string {
	str := ""
	for index, value := range args {
		if index == len(args)-1 {
			str += value
			break
		}
		str += value + "\n"
	}
	return str
}
