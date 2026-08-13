
package piscine

import "fmt"

func UltimateDivMod(a *int, b *int) {
	cup := *a % *b 
	*a = *a / *b
	*b = cup % *a
}

