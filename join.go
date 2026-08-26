package main

import "fmt"

func Join(strs []string, sep string) string {
	str := ""
	for index, value := range strs {
		if index != len(strs)-1 {
			str += value + sep
		} else {
			str += value
		}
	}

	return str
}

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, "-"))
}
