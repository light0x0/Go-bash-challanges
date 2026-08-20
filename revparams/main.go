package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	names := os.Args
	for index := len(names) - 1; index > 0; index-- {
		name := names[index]
		for indexa := 0; indexa < len(name); indexa++ {
			z01.PrintRune(rune(names[index][indexa]))
		}
		z01.PrintRune('\n')
	}
}
