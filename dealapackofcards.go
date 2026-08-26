package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	total := 0
	for range deck {
		total++
	}
	math := total / 4
	index := math

	for player := 1; player <= 4; player++ {
		fmt.Printf("Player %d: ", player)

		for indexa := index - math; indexa < index; indexa++ {
			if indexa > index-math {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			fmt.Printf("%d", deck[indexa])
		}

		z01.PrintRune('\n')
		index += math
	}
}
