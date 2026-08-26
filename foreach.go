package piscine

func PrintNbr(n int) {
	fmt.Print(n)
}

func ForEach(f func(int), a []int) {
	for _, value := range a {
		f(value)
	}
}
