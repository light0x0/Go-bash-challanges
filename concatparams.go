package piscine

func ConcatParams(args []string) string {
	str := ""
	for _, value := range args {
		str += value + "\n"
	}
	return str
}
