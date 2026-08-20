package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	step := 0
	for index, value := range name {
		if value == '/' {
			step = index + 1
		}
	}

	for index := step; index < len(name); index++ {
		z01.PrintRune(rune(name[index]))
	}
}
