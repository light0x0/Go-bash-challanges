package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	names := os.Args

	for index := 1; index < len(names); index++ {
		for indexa := index + 1; indexa < len(names); indexa++ {
			if names[index] > names[indexa] {
				names[index], names[indexa] = names[indexa], names[index]
			}
		}
	}

	for index := 1; index < len(names); index++ {
		for indexa := 0; indexa < len(names[index]); indexa++ {
			z01.PrintRune(rune(names[index][indexa]))
		}
		z01.PrintRune('\n')
	}
}
