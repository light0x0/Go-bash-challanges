package piscine

func Swap(a *int, b *int) {
	cup := *a
	*a = *b
	*a = cup
}
