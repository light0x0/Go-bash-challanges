package piscine

func JumpOver(str string) string {
	if str == "" || len(str) == 2 {
		return "\n"
	}

	s := ""

	for index := 2; index < len(str); index += 3 {
		s += string(str[index])
	}
	return s + "\n"
}
