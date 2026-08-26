package piscine

// import "fmt"

func Unmatch(a []int) int {
	for index := 0; index < len(a); index++ {
		match := 0
		for index2 := 0; index2 < len(a); index2++ {
			//fmt.Println(a[index],a[index2])
			if a[index] == a[index2] {
				match++
			}
		}
		if match%2 != 0 {
			return a[index]
		}
	}

	return -1
}
