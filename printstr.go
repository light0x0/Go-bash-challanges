package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	convert := []rune(s)
	for index := 0; index < len(convert); index++ {
		z01.PrintRune(convert[index])
	}
}
