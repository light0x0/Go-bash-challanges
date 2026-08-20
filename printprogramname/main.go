package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, value := range os.Args[0][:2] {
		z01.PrintRune(value)
	}
}
