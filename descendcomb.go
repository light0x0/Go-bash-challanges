
package main
import "github.com/01-edu/z01"

func DescendComb() {

	for v1 := 9; v1 >= 0; v1-- {
		for v2 := 9; v2 >= 1; v2-- {
			for v3 := 9; v3 >= 0; v3-- {
				for v4 := 8; v4 >= 0; v4-- {
					if (v1 > v3) || (v1 == v3 && v2 > v4){
						z01.PrintRune(rune('0' + v1))
						z01.PrintRune(rune('0' + v2))
						z01.PrintRune(' ')
						z01.PrintRune(rune('0' + v3))
						z01.PrintRune(rune('0' + v4))
						if !(v1 ==  0&& v2 == 1 && v3 == 0 && v4 == 0){
							z01.PrintRune(',')
						}
					}
				}
			}
		}
	}
}

func main() {
	DescendComb()
}